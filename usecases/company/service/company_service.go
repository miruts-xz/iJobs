package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
)

// CompanyServiceImpl implements CompanyService interface
type CompanyServiceImpl struct {
	cpRepo company.CompanyRepository
}

// NewCompanyServiceImpl returns new CompanyServiceImpl
func (cps *CompanyServiceImpl) NewCompanyServiceImpl(cpr company.CompanyRepository) *CompanyServiceImpl {
	return &CompanyServiceImpl{cpRepo: cpr}
}

// Company return all companys
func (cps *CompanyServiceImpl) Companys() ([]entity.Company, error) {
	return cps.cpRepo.Companys()
}

// Company return company with a given id
func (cps *CompanyServiceImpl) Company(id int) (entity.Company, error) {
	return cps.cpRepo.Company(id)
}

// UpdateCompany updates a given company
func (cps *CompanyServiceImpl) UpdateCompany(cp entity.Company) error {
	return cps.cpRepo.UpdateCompany(cp)
}

// DeleteCompany deletes company with a given id
func (cps *CompanyServiceImpl) DeleteCompany(id int) error {
	return cps.cpRepo.DeleteCompany(id)
}

// StoreCompany stores new Company
func (cps *CompanyServiceImpl) StoreCompany(cp entity.Company) error {
	return cps.cpRepo.StoreCompany(cp)
}
