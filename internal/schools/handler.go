package schools

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"schools/internal"
)

type SchoolHandler struct {
	schoolService *SchoolService
}

func NewSchoolHandler(schoolService *SchoolService) internal.HandlerInterface {
	return &SchoolHandler{
		schoolService: schoolService,
	}
}

func (h *SchoolHandler) Register(e *echo.Echo) {
	e.POST("/schools", h.CreateNewSchool)
	e.GET("/schools", h.ListSchools)
}

func (h *SchoolHandler) CreateNewSchool(c echo.Context) error {
	var school School
	if err := c.Bind(&school); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdSchool, err := h.schoolService.CreateSchool(school.Name, school.Address)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create school")
	}

	return c.JSON(http.StatusCreated, createdSchool.ID)
}

func (h *SchoolHandler) ListSchools(c echo.Context) error {
	schools, err := h.schoolService.ListSchools()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to load schools")
	}

	return c.JSON(http.StatusOK, schools)
}
