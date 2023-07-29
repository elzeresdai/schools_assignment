package schools_students

import (
	"fmt"
	"schools/internal/schools"
	"schools/internal/students"
)

type SchoolStudentService struct {
	schoolStudentRepo SchoolStudentRepositoryInterface
	studentRepo       students.StudentRepositoryInterface
	schoolRepo        schools.SchoolRepositoryInterface
}

func NewSchoolStudentService(schoolStudentRepo SchoolStudentRepositoryInterface, studentRepo students.StudentRepositoryInterface,
	schoolRepo schools.SchoolRepositoryInterface) *SchoolStudentService {
	return &SchoolStudentService{
		schoolStudentRepo: schoolStudentRepo,
		studentRepo:       studentRepo,
		schoolRepo:        schoolRepo,
	}
}

func (s *SchoolStudentService) ListAllStudentsInOrder() ([]*students.Student, error) {
	return s.schoolStudentRepo.ListAllStudentsInOrder()
}

func (s *SchoolStudentService) AddStudentToSchool(schoolID string, student students.Student) (*students.Student, error) {
	// Check if the school exists
	if _, err := s.schoolRepo.GetSchoolByID(schoolID); err != nil {
		return nil, fmt.Errorf("school not found: %w", err)
	}

	// Check if the student exists
	if _, err := s.studentRepo.GetStudentByID(student.ID); err != nil {
		return nil, fmt.Errorf("student not found: %w", err)
	}

	err := s.schoolStudentRepo.AddStudentToSchool(schoolID, student.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to add student to school: %w", err)
	}

	return &student, nil
}

func (s *SchoolStudentService) GetStudentsBySchoolID(schoolID string) ([]*students.Student, error) {
	// Check if the school exists
	if _, err := s.schoolRepo.GetSchoolByID(schoolID); err != nil {
		return nil, fmt.Errorf("school not found: %w", err)
	}

	return s.schoolStudentRepo.GetStudentsBySchoolID(schoolID)
}
