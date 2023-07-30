package schools

import (
	"errors"
	"testing"
)

type mockSchoolRepository struct {
	schools map[string]*School
}

func (m *mockSchoolRepository) CreateSchool(school *School) error {
	if m.schools == nil {
		m.schools = make(map[string]*School)
	}

	if _, exists := m.schools[school.ID]; exists {
		return errors.New("school already exists")
	}

	m.schools[school.ID] = school
	return nil
}

func (m *mockSchoolRepository) GetAllSchools() ([]*School, error) {
	var allSchools []*School
	for _, school := range m.schools {
		allSchools = append(allSchools, school)
	}
	return allSchools, nil
}
func (m *mockSchoolRepository) GetSchoolByID(id string) (*School, error) {
	school, exists := m.schools[id]
	if !exists {
		return nil, errors.New("school not found")
	}
	return school, nil
}

func TestSchoolService(t *testing.T) {
	repo := &mockSchoolRepository{}
	service := NewSchoolService(repo)

	// Test CreateSchool
	school := &School{
		Name:    "Test School",
		Address: "Test Address",
	}
	_, err := service.CreateSchool(school.Name, school.Address)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = service.CreateSchool(school.Name, school.Address)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Test ListSchools
	schools, err := service.ListSchools()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(schools) != 1 {
		t.Errorf("Expected 1 school, got %d", len(schools))
	}

	if schools[0].Name != school.Name || schools[0].Address != school.Address {
		t.Errorf("Expected school name and address to match, got %v", schools[0])
	}
}
