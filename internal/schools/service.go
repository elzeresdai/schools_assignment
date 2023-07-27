package schools

type SchoolService struct {
	schoolRepo SchoolRepositoryInterface
}

// NewSchoolService creates a new instance of SchoolService
func NewSchoolService(schoolRepo SchoolRepositoryInterface) *SchoolService {
	return &SchoolService{
		schoolRepo: schoolRepo,
	}
}

// CreateSchool creates a new school
func (s *SchoolService) CreateSchool(name, address string) (*School, error) {
	school := &School{
		Name:    name,
		Address: address,
	}

	err := s.schoolRepo.CreateSchool(school)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// ListSchools returns a list of all schools
func (s *SchoolService) ListSchools() ([]*School, error) {
	return s.schoolRepo.GetAllSchools()
}
