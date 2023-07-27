package students

type StudentRepositoryInterface interface {
	CreateStudent(student *Student) error
	GetStudentByID(id string) (*Student, error)
}
