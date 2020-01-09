package service

import (
	entity2 "github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
)

type CompanyServiceImpl struct {
	compRepo company.CompanyRepository
}

func NewCompanyServiceImpl(cr company.CompanyRepository) *CompanyServiceImpl {
	return &CompanyServiceImpl{compRepo: cr}
}
func (cs *CompanyServiceImpl) Companies() ([]entity2.Company, error) {
	return cs.compRepo.Companies()
}
func (cs *CompanyServiceImpl) Company(cid int) (entity2.Company, error) {
	return cs.compRepo.Company(cid)
}
func (cs *CompanyServiceImpl) UpdateCompany(cmp *entity2.Company) (*entity2.Company, error) {
	return cs.compRepo.UpdateCompany(cmp)
}
func (cs *CompanyServiceImpl) DeleteCompany(cid int) (entity2.Company, error) {
	return cs.compRepo.DeleteCompany(cid)
}
func (cs *CompanyServiceImpl) StoreCompany(cmp *entity2.Company) (*entity2.Company, error) {
	return cs.compRepo.StoreCompany(cmp)
}
func (cs *CompanyServiceImpl) PostedJobs(cid int) ([]entity2.Job, error) {
	return cs.compRepo.PostedJobs(cid)
}
