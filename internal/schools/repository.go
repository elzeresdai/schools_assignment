package schools

import (
	"errors"
	"github.com/google/uuid"
	"sync"
)

type SchoolRepository struct {
	schools map[string]*School
	mu      sync.RWMutex
}

func NewSchoolRepository() SchoolRepositoryInterface {
	return &SchoolRepository{
		schools: make(map[string]*School),
	}
}

func (r *SchoolRepository) CreateSchool(school *School) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New().String()
	school.ID = id
	r.schools[id] = school

	return nil
}

func (r *SchoolRepository) GetAllSchools() ([]*School, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var schools []*School
	for _, school := range r.schools {
		schools = append(schools, school)
	}
	return schools, nil
}

func (r *SchoolRepository) GetSchoolByID(schoolId string) (*School, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	school, ok := r.schools[schoolId]
	if !ok {
		return nil, errors.New("school not found")
	}
	return school, nil
}
