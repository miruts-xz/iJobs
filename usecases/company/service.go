package company

import (
	entity2 "github.com/miruts/iJobs/entity"
)

type CompanyService interface {
	Companies() ([]entity2.Company, error)
	Company(cid int) (entity2.Company, error)
	UpdateCompany(cmp *entity2.Company) (*entity2.Company, error)
	DeleteCompany(cid int) (entity2.Company, error)
	StoreCompany(cmp *entity2.Company) (*entity2.Company, error)
	PostedJobs(cid int) ([]entity2.Job, error)
}
