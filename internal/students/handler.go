package students

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"schools/internal"
)

type StudentHandler struct {
	studentService *StudentService
}

func NewStudentHandler(studentService *StudentService) internal.HandlerInterface {
	return &StudentHandler{
		studentService: studentService,
	}
}

func (h *StudentHandler) Register(e *echo.Echo) {
	e.POST("/student", h.CreateStudent)
}

func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var student Student
	if err := c.Bind(&student); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdStudent, err := h.studentService.CreateStudent(student.Name, student.AverageGrade)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to create student")
	}

	return c.JSON(http.StatusCreated, createdStudent.ID)
}
