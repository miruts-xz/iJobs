package company

import entity "github.com/miruts/iJobs/entity/gorm-entity"

type CompanyRepository interface {
	Companies() ([]entity.Company, error)
	Company(cid int) (entity.Company, error)
	UpdateCompany(cmp *entity.Company) (*entity.Company, error)
	DeleteCompany(cid int) (entity.Company, error)
	StoreCompany(cmp *entity.Company) (*entity.Company, error)
	PostedJobs(cid int) ([]entity.Job, error)
}
