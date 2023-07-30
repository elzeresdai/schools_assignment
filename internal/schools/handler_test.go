package schools

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSchoolHandler_CreateNewSchool(t *testing.T) {
	// Create a new instance of SchoolService with a mock repository
	schoolService := NewSchoolService(&mockSchoolRepository{})
	handler := NewSchoolHandler(schoolService)

	e := echo.New()
	handler.Register(e)

	payload := map[string]string{
		"Name":    "Test School",
		"Address": "Test Address",
	}
	jsonPayload, _ := json.Marshal(payload)

	// Create a new HTTP request to the "/schools" endpoint with POST method
	req := httptest.NewRequest(http.MethodPost, "/schools", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Handle the request using the Echo instance
	e.ServeHTTP(rec, req)

	// Check if the response code is as expected
	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

	responseID := rec.Body.String()
	if responseID == "" {
		t.Error("expected non-empty school ID, got empty")
	}
}

func TestSchoolHandler_ListSchools(t *testing.T) {
	// Create a new instance of SchoolService with a mock repository
	schoolService := NewSchoolService(&mockSchoolRepository{})
	handler := NewSchoolHandler(schoolService)

	e := echo.New()
	handler.Register(e)

	// Create a new HTTP request to the "/schools" endpoint with GET method
	req := httptest.NewRequest(http.MethodGet, "/schools", nil)

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Handle the request using the Echo instance
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var schools []*School
	if err := json.Unmarshal(rec.Body.Bytes(), &schools); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	expectedNumSchools := 0 // Set the expected number of schools
	if len(schools) != expectedNumSchools {
		t.Errorf("expected %d schools, got %d", expectedNumSchools, len(schools))
	}
}
