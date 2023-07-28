package schools_students

import "schools/internal/students"

type SchoolStudentRepositoryInterface interface {
	AddStudentToSchool(schoolID, studentID string) error
	GetStudentsBySchoolID(schoolID string) ([]*students.Student, error)
}
