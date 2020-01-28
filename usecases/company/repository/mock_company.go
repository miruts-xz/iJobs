package repository

import (
	"github.com/miruts/iJobs/entity"
)

type CompanyMockRepository struct {
}

func NewCompanyMockRepository() *CompanyMockRepository {
	return &CompanyMockRepository{}
}

func (c *CompanyMockRepository) Companies() ([]entity.Company, error) {
	panic("implement me")
}

func (c *CompanyMockRepository) Company(cid int) (entity.Company, error) {
	panic("implement me")
}

func (c *CompanyMockRepository) UpdateCompany(cmp *entity.Company) (*entity.Company, error) {
	return cmp, nil
}

func (c *CompanyMockRepository) DeleteCompany(cid int) (entity.Company, error) {
	panic("implement me")
}

func (c *CompanyMockRepository) StoreCompany(cmp *entity.Company) (*entity.Company, error) {
	panic("implement me")
}

func (c *CompanyMockRepository) PostedJobs(cid int) ([]entity.Job, error) {
	panic("implement me")
}

func (c *CompanyMockRepository) CompanyByEmail(email string) (entity.Company, error) {
	if email == entity.Companymock1.Email {
		return entity.Companymock1, nil
	} else if email == entity.Companymock2.Email {
		return entity.Companymock2, nil
	}
	return entity.Company{}, nil
}

func (c *CompanyMockRepository) CompanyAddress(id uint) (entity.Address, error) {
	panic("implement me")
}

func (c *CompanyMockRepository) UserRoles(user *entity.Company) ([]entity.Role, []error) {
	panic("implement me")
}

func (c *CompanyMockRepository) EmailExists(email string) bool {
	panic("implement me")
}

func (c *CompanyMockRepository) UsernameExists(email string) bool {
	panic("implement me")
}

func (c *CompanyMockRepository) PhoneExists(phone string) bool {
	panic("implement me")
}

func (c *CompanyMockRepository) JobExists(cm_id int, job string) bool {
	if entity.Jobmock1.CompanyID == uint(cm_id) {
		return true
	} else if entity.Jobmock2.CompanyID == uint(cm_id) {
		return true
	}
	return false
}
