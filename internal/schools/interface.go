package schools

type SchoolRepositoryInterface interface {
	CreateSchool(school *School) error
	GetAllSchools() ([]*School, error)
	GetSchoolByID(id string) (*School, error)
}
