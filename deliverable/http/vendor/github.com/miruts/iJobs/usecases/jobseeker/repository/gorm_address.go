package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

// JobseekerGormRepositoryIMpl implements JobseekerRepository interface
type AddressGormRepositoryImpl struct {
	conn *gorm.DB
}

// NewJobseekerRepositoryImpl returns new JobseekerGormRepositoryIMpl
func NewAddressGormRepositoryImpl(jsr *gorm.DB) *AddressGormRepositoryImpl {
	return &AddressGormRepositoryImpl{conn: jsr}
}

// JobSeekers retrieves and returns all jobseekers
func (jsr *AddressGormRepositoryImpl) Addresses() ([]entity.Address, error) {
	var addresses []entity.Address
	errs := jsr.conn.Find(&addresses).GetErrors()
	if errs != nil {
		fmt.Printf("Error: %v", errs)
		return addresses, errs[0]
	}
	return addresses, nil
}

// JobSeeker return a jobseeker with given id
func (jsr *AddressGormRepositoryImpl) Address(id int) (entity.Address, error) {
	var addresss entity.Address
	errs := jsr.conn.First(&addresss, id).GetErrors()
	if len(errs) > 0 {
		return addresss, errs[0]
	}
	return addresss, nil
}

// UpdateJobSeeker updates a given jobseeker
func (jsr *AddressGormRepositoryImpl) UpdateAddress(adr *entity.Address) (*entity.Address, error) {
	address := adr
	errs := jsr.conn.Save(&address).GetErrors()
	if len(errs) > 0 {
		return address, errs[0]
	}
	return address, nil
}

// DeleteJobSeeker deletes a jobseeker with a given id
func (jsr *AddressGormRepositoryImpl) DeleteAddress(id int) (entity.Address, error) {
	address, err := jsr.Address(id)
	if err != nil {
		return address, err
	}
	errs := jsr.conn.Delete(address, id).GetErrors()
	if len(errs) > 0 {
		return address, errs[0]
	}
	return address, nil
}

// StoreJobSeeker stores new jobseeker
func (jsr *AddressGormRepositoryImpl) StoreAddress(adr *entity.Address) (*entity.Address, error) {
	address := adr
	errs := jsr.conn.Create(address).GetErrors()
	if len(errs) > 0 {
		return address, errs[0]
	}
	return address, nil
}
