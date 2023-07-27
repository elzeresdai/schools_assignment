package schools

import (
	"errors"
	"github.com/google/uuid"
	"schools/internal/students"
	"sync"
)

type SchoolRepository struct {
	schools      map[string]*School
	schoolsMutex sync.Mutex
}

func NewSchoolRepository() SchoolRepositoryInterface {
	return &SchoolRepository{
		schools: make(map[string]*School),
	}
}

func (r *SchoolRepository) CreateSchool(school *School) error {
	id := uuid.New().String()
	school.ID = id
	r.schools[id] = school
	return nil
}

func (r *SchoolRepository) GetAllSchools() ([]*School, error) {
	var schools []*School
	for _, school := range r.schools {
		schools = append(schools, school)
	}
	return schools, nil
}

func (r *SchoolRepository) GetSchoolByID(schoolId string) (*School, error) {
	school, ok := r.schools[schoolId]
	if !ok {
		return nil, errors.New("school not found")
	}
	return school, nil
}

func (r *SchoolRepository) AddStudentToSchool(schoolID string, student *students.Student) error {
	//TODO implement me
	panic("implement me")
}

func (r *SchoolRepository) GetSchoolStudents(schoolID string) ([]students.Student, error) {
	//TODO implement me
	panic("implement me")
}
