package students

type StudentService struct {
	studentsRepo StudentRepositoryInterface
}

// NewStudentService NewSchoolService NewStudentService creates a new instance of StudentService
func NewStudentService(schoolRepo StudentRepositoryInterface) *StudentService {
	return &StudentService{
		studentsRepo: schoolRepo,
	}
}

func (s *StudentService) CreateStudent(name string, averageGrade float64) (*Student, error) {
	student := &Student{
		Name:         name,
		AverageGrade: averageGrade,
	}

	err := s.studentsRepo.CreateStudent(student)

	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *StudentService) GetStudentById(id string) (*Student, error) {
	student, err := s.studentsRepo.GetStudentByID(id)

	if err != nil {
		return nil, err
	}

	return student, nil
}
