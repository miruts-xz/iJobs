package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
)

type CompanyServiceImpl struct {
	compRepo company.CompanyRepository
}

func NewCompanyServiceImpl(cr company.CompanyRepository) *CompanyServiceImpl {
	return &CompanyServiceImpl{compRepo: cr}
}
func (cs *CompanyServiceImpl) Companies() ([]entity.Company, error) {
	return cs.compRepo.Companies()
}
func (cs *CompanyServiceImpl) Company(cid int) (entity.Company, error) {
	return cs.compRepo.Company(cid)
}
func (cs *CompanyServiceImpl) UpdateCompany(cmp *entity.Company) (*entity.Company, error) {
	return cs.compRepo.UpdateCompany(cmp)
}
func (cs *CompanyServiceImpl) DeleteCompany(cid int) (entity.Company, error) {
	return cs.compRepo.DeleteCompany(cid)
}
func (cs *CompanyServiceImpl) StoreCompany(cmp *entity.Company) (*entity.Company, error) {
	return cs.compRepo.StoreCompany(cmp)
}
func (cs *CompanyServiceImpl) PostedJobs(cid int) ([]entity.Job, error) {
	return cs.compRepo.PostedJobs(cid)
}
func (cs *CompanyServiceImpl) CompanyByEmail(email string) (entity.Company, error) {
	return cs.compRepo.CompanyByEmail(email)
}
func (cs *CompanyServiceImpl) CompanyAddress(id uint) (entity.Address, error) {
	return cs.compRepo.CompanyAddress(id)
}
