package schools_students

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"schools/internal"
	"schools/internal/students"
)

type Handler struct {
	service *SchoolStudentService
}

func NewHandler(service *SchoolStudentService) internal.HandlerInterface {
	return &Handler{
		service: service,
	}
}

func (h Handler) Register(e *echo.Echo) {
	e.POST("/schools/:schoolID/students", h.AddStudentSchool)
	e.GET("/schools/:schoolID/students", h.GetStudentsBySchoolID)
	e.GET("/students/all", h.ListAllStudentsInOrder)
}

func (h *Handler) AddStudentSchool(c echo.Context) error {
	schoolID := c.Param("schoolID")

	var student students.Student
	if err := c.Bind(&student); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdStudent, err := h.service.AddStudentToSchool(schoolID, student)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to add student to school")
	}

	return c.JSON(http.StatusCreated, createdStudent)
}

func (h *Handler) ListAllStudentsInOrder(c echo.Context) error {
	students, err := h.service.ListAllStudentsInOrder()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to load students")
	}

	return c.JSON(http.StatusOK, students)
}

func (h *Handler) GetStudentsBySchoolID(c echo.Context) error {
	schoolID := c.Param("schoolID")
	students, err := h.service.GetStudentsBySchoolID(schoolID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to load students")
	}
	return c.JSON(http.StatusOK, students)
}
