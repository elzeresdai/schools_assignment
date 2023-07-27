package schools

import "schools/internal/students"

type SchoolRepositoryInterface interface {
	CreateSchool(school *School) error
	GetAllSchools() ([]*School, error)
	GetSchoolByID(id string) (*School, error)
	AddStudentToSchool(schoolID string, student *students.Student) error
	GetSchoolStudents(schoolID string) ([]students.Student, error)
}
