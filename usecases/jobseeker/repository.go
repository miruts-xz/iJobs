package jobseeker

import (
	"github.com/miruts/iJobs/entity"
)

// JobseekerRepository interface defines all jobseeker related data/database operations
type JobseekerRepository interface {
	JobSeekers() ([]entity.Jobseeker, error)
	JobSeeker(id int) (entity.Jobseeker, error)
	UpdateJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error)
	DeleteJobSeeker(id int) (entity.Jobseeker, error)
	StoreJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error)
	JsCategories(id int) ([]entity.Category, error)
	RemoveIntCategory(jsid, jcid int) error
	AddIntCategory(jsid, jcid int) error
	SetAddress(jsid, addid int) error
	JobseekerByEmail(email string) (entity.Jobseeker, error)
	JobseekerByUsername(uname string) (entity.Jobseeker, error)
	ApplicationJobseeker(id int) (entity.Jobseeker, error)
	UserRoles(user *entity.Jobseeker) ([]entity.Role, []error)
	PhoneExists(phone string) bool
	EmailExists(email string) bool
	UsernameExists(email string) bool
	AlreadyApplied(id uint, id2 uint) bool
}

// AddressRepository interface defines all jobseeker related data/database operations
type AddressRepository interface {
	Addresses() ([]entity.Address, error)
	Address(id int) (entity.Address, error)
	UpdateAddress(category *entity.Address) (*entity.Address, error)
	DeleteAddress(id int) (entity.Address, error)
	StoreAddress(category *entity.Address) (*entity.Address, error)
}
