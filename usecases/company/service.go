package company

import "github.com/miruts/iJobs/entity"

// CompanyService interface that defines Company related services
type JobseekerService interface {
	Companys() ([]entity.Company, error)
	Company(id int) (entity.Company, error)
	UpdateCompany(cp entity.Company) error
	DeleteCompany(id int) error
	StoreCompany(cp entity.Company) error
}
