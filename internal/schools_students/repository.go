package schools_students

import "sync"

type SchoolStudentRepository struct {
	schools      map[string]*SchoolStudents
	schoolsMutex sync.RWMutex
}
