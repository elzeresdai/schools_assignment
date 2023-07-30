package students

import (
	"errors"
	"github.com/google/uuid"
	"sort"
	"sync"
)

type StudentRepository struct {
	students map[string]*Student
	mu       sync.RWMutex
}

func NewStudentRepository() StudentRepositoryInterface {
	return &StudentRepository{
		students: make(map[string]*Student),
	}
}

// CreateStudent AddStudent adds a new student to the repository
func (r *StudentRepository) CreateStudent(student *Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New().String()
	student.ID = id
	r.students[id] = student

	return nil
}

func (r *StudentRepository) GetStudentByID(id string) (*Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	student, ok := r.students[id]
	if !ok {
		return nil, errors.New("student not found")
	}

	return student, nil
}

func (r *StudentRepository) GetHighestGradedStudents(n int) ([]*Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	grades := make([]struct {
		StudentID    string
		AverageGrade float64
	}, 0, len(r.students))

	for studentID, student := range r.students {
		grades = append(grades, struct {
			StudentID    string
			AverageGrade float64
		}{StudentID: studentID, AverageGrade: student.AverageGrade})
	}

	sort.Slice(grades, func(i, j int) bool {
		return grades[i].AverageGrade > grades[j].AverageGrade
	})

	highestGradedStudents := make([]*Student, 0, n)
	for i := 0; i < n && i < len(grades); i++ {
		student, ok := r.students[grades[i].StudentID]
		if ok {
			highestGradedStudents = append(highestGradedStudents, student)
		}
	}

	return highestGradedStudents, nil
}
