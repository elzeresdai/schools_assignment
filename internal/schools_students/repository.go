package schools_students

import (
	"schools/internal/students"
	"sort"
	"sync"
)

type SchoolStudentRepository struct {
	schoolStudents []*SchoolStudents
	studentRepo    students.StudentRepositoryInterface
	mu             sync.RWMutex
	counter        int64
}

func NewSchoolStudentRepository(studentRepo students.StudentRepositoryInterface) SchoolStudentRepositoryInterface {
	return &SchoolStudentRepository{
		schoolStudents: make([]*SchoolStudents, 0),
		studentRepo:    studentRepo,
		counter:        0,
	}
}

func (r *SchoolStudentRepository) ListAllStudentsInOrder() ([]*students.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	studentsByID := make(map[string]*students.Student)

	for _, schoolStudent := range r.schoolStudents {
		// Check if the student is already added
		if _, ok := studentsByID[schoolStudent.StudentID]; !ok {
			student, err := r.studentRepo.GetStudentByID(schoolStudent.StudentID)
			if err != nil {
				return nil, err
			}
			studentsByID[schoolStudent.StudentID] = student
		}
	}

	allStudents := make([]*students.Student, 0, len(r.schoolStudents))
	for _, schoolStudent := range r.schoolStudents {
		student, ok := studentsByID[schoolStudent.StudentID]
		if ok {
			allStudents = append(allStudents, student)
		}
	}

	// Sort the students by orderId
	sort.Slice(allStudents, func(i, j int) bool {
		return r.getStudentOrder(allStudents[i].ID) < r.getStudentOrder(allStudents[j].ID)
	})

	return allStudents, nil
}

func (r *SchoolStudentRepository) getStudentOrder(studentID string) int64 {
	for _, schoolStudent := range r.schoolStudents {
		if schoolStudent.StudentID == studentID {
			return schoolStudent.OrderID
		}
	}
	return -1 // Return -1 if the student is not exist
}

func (r *SchoolStudentRepository) AddStudentToSchool(schoolID, studentID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	record := &SchoolStudents{
		SchoolID:  schoolID,
		StudentID: studentID,
		OrderID:   r.getOrderID(),
	}

	r.schoolStudents = append(r.schoolStudents, record)

	return nil
}

func (r *SchoolStudentRepository) getOrderID() int64 {
	r.counter++
	return r.counter
}

func (r *SchoolStudentRepository) GetStudentsBySchoolID(schoolID string) ([]*students.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	studentsBySchool := make([]*students.Student, 0)

	for _, schoolStudent := range r.schoolStudents {
		if schoolStudent.SchoolID == schoolID {
			student, err := r.studentRepo.GetStudentByID(schoolStudent.StudentID)
			if err != nil {
				return nil, err
			}
			studentsBySchool = append(studentsBySchool, student)
		}
	}

	return studentsBySchool, nil
}
