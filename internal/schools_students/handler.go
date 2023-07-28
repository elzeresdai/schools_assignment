package schools_students

import (
	"github.com/labstack/echo/v4"
	"schools/internal"
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
}

func (h *Handler) AddStudentSchool(c echo.Context) error {
	return nil
}
