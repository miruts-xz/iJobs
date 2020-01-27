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

// Companies retrieves and returns all companies
func (cs *CompanyServiceImpl) Companies() ([]entity.Company, error) {
	return cs.compRepo.Companies()
}

// Company return a Company with given id
func (cs *CompanyServiceImpl) Company(cid int) (entity.Company, error) {
	return cs.compRepo.Company(cid)
}

// UpdateCompany updates a given company
func (cs *CompanyServiceImpl) UpdateCompany(cmp *entity.Company) (*entity.Company, error) {
	return cs.compRepo.UpdateCompany(cmp)
}

// DeleteCompany deletes a company with a given id
func (cs *CompanyServiceImpl) DeleteCompany(cid int) (entity.Company, error) {
	return cs.compRepo.DeleteCompany(cid)
}

// StoreCompany stores new company
func (cs *CompanyServiceImpl) StoreCompany(cmp *entity.Company) (*entity.Company, error) {
	return cs.compRepo.StoreCompany(cmp)
}

// Posted jobs retrieves jobs jobs posted by company
func (cs *CompanyServiceImpl) PostedJobs(cid int) ([]entity.Job, error) {
	return cs.compRepo.PostedJobs(cid)
}

// CompanyByEmail retrieves company given email
func (cs *CompanyServiceImpl) CompanyByEmail(email string) (entity.Company, error) {
	return cs.compRepo.CompanyByEmail(email)
}

// CompanyAddress retrieves address of a company given company id
func (cs *CompanyServiceImpl) CompanyAddress(id uint) (entity.Address, error) {
	return cs.compRepo.CompanyAddress(id)
}

// UserRoles returns list of roles a user has
func (jss *CompanyServiceImpl) UserRoles(user *entity.Company) ([]entity.Role, []error) {
	userRoles, errs := jss.compRepo.UserRoles(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}

// PhoneExists check if there is a user with a given phone number
func (jss *CompanyServiceImpl) PhoneExists(phone string) bool {
	exists := jss.compRepo.PhoneExists(phone)
	return exists
}
func (jss *CompanyServiceImpl) UsernameExists(email string) bool {
	exists := jss.compRepo.UsernameExists(email)
	return exists
}

// EmailExists checks if there exist a user with a given email address
func (jss *CompanyServiceImpl) EmailExists(email string) bool {
	exists := jss.compRepo.EmailExists(email)
	return exists
}
func (jss *CompanyServiceImpl) JobExists(cm_id int, job string) bool {
	exists := jss.compRepo.JobExists(cm_id, job)
	return exists
}
