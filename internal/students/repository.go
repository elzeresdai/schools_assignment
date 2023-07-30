package students

import (
	"errors"
	"github.com/google/uuid"
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
	//TODO implement me
	panic("implement me")
}
