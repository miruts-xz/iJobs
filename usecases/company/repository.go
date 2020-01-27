package company

import (
	"github.com/miruts/iJobs/entity"
)

// CompanyRepository interface represents all company data repository actions
type CompanyRepository interface {
	Companies() ([]entity.Company, error)
	Company(cid int) (entity.Company, error)
	UpdateCompany(cmp *entity.Company) (*entity.Company, error)
	DeleteCompany(cid int) (entity.Company, error)
	StoreCompany(cmp *entity.Company) (*entity.Company, error)
	PostedJobs(cid int) ([]entity.Job, error)
	CompanyByEmail(email string) (entity.Company, error)
	CompanyAddress(id uint) (entity.Address, error)
	UserRoles(user *entity.Company) ([]entity.Role, []error)
	EmailExists(email string) bool
	UsernameExists(email string) bool
	PhoneExists(phone string) bool
	JobExists(cm_id int, job string) bool
}
