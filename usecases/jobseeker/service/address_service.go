package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

type AddressServiceImpl struct {
	addRepo jobseeker.AddressRepository
}

// NewJobseekerServiceImpl returns new JobseekerServiceImpl
func NewAddressServiceImpl(jsr jobseeker.AddressRepository) *AddressServiceImpl {
	return &AddressServiceImpl{addRepo: jsr}
}

// JobSeekers return all jobseekers
func (jss *AddressServiceImpl) Addresses() ([]entity.Address, error) {
	return jss.addRepo.Addresses()
}

// JobSeeker return jobseeker with a given id
func (jss *AddressServiceImpl) Address(id int) (entity.Address, error) {
	return jss.addRepo.Address(id)
}

// UpdateJobSeeker updates a given jobseeker
func (jss *AddressServiceImpl) UpdateAddress(js *entity.Address) (*entity.Address, error) {
	return jss.addRepo.UpdateAddress(js)
}

// DeleteJobSeeker deletes jobseeker with a given id
func (jss *AddressServiceImpl) DeleteAddress(id int) (entity.Address, error) {
	return jss.addRepo.DeleteAddress(id)
}

// StoreJobSeeker stores new jobseeker
func (jss *AddressServiceImpl) StoreAddress(js *entity.Address) (*entity.Address, error) {
	return jss.addRepo.StoreAddress(js)
}
