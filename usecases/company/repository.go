package company

import "github.com/miruts/iJobs/entity"

// CompanyRepository interface defines all Company related data/database operations
type CompanyRepository interface {
	Companys() ([]entity.Company, error)
	Company(id int) (entity.Company, error)
	UpdateCompany(cp entity.Company) error
	DeleteCompany(id int) error
	StoreCompany(cp entity.Company) error
}
