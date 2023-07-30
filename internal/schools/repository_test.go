package schools

import (
	"testing"
)

func TestSchoolRepository(t *testing.T) {
	// Create a new instance of the repository
	repo := NewSchoolRepository()

	// Create a test school
	testSchool := &School{
		Name:    "Test School",
		Address: "Test Address",
	}

	// Test CreateSchool method
	err := repo.CreateSchool(testSchool)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test GetSchoolByID method
	school, err := repo.GetSchoolByID(testSchool.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if school == nil {
		t.Errorf("Expected school to be retrieved, got nil")
	}
	if school.Name != testSchool.Name || school.Address != testSchool.Address {
		t.Errorf("Expected school data to match, got %v", school)
	}

	// Test GetAllSchools method
	schools, err := repo.GetAllSchools()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(schools) != 1 {
		t.Errorf("Expected 1 school, got %d", len(schools))
	}

	// Test non-existing school
	_, err = repo.GetSchoolByID("non_existing_id")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
