package jobseeker

import (
	"github.com/miruts/iJobs/entity"
)

// JobseekerService interface that defines jobseeker related services
type JobseekerService interface {
	JobSeekers() ([]entity.Jobseeker, error)
	JobSeeker(id int) (entity.Jobseeker, error)
	UpdateJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error)
	DeleteJobSeeker(id int) (entity.Jobseeker, error)
	StoreJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error)
	Suggestions(id int) ([]entity.Job, error)
	AddIntCategory(jsid, jcid int) error
	RemoveIntCategory(jsid, jcid int) error
	SetAddress(jsid, addid int) error
}
type AddressService interface {
	Addresses() ([]entity.Address, error)
	Address(id int) (entity.Address, error)
	UpdateAddress(category *entity.Address) (*entity.Address, error)
	DeleteAddress(id int) (entity.Address, error)
	StoreAddress(category *entity.Address) (*entity.Address, error)
}
