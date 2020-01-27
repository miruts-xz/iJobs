package company

import (
	"github.com/miruts/iJobs/entity"
)

// CompanyService interface represents all company services
type CompanyService interface {
	Companies() ([]entity.Company, error)
	Company(cid int) (entity.Company, error)
	UpdateCompany(cmp *entity.Company) (*entity.Company, error)
	DeleteCompany(cid int) (entity.Company, error)
	StoreCompany(cmp *entity.Company) (*entity.Company, error)
	PostedJobs(cid int) ([]entity.Job, error)
	CompanyByEmail(email string) (entity.Company, error)
	CompanyAddress(id uint) (entity.Address, error)
	UserRoles(user *entity.Company) ([]entity.Role, []error)
	UsernameExists(username string) bool
	JobExists(cm_id int, job string) bool
}
