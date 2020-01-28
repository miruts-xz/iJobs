package repository

import (
	"errors"
	"github.com/miruts/iJobs/entity"
)

type JobseekerMockRepository struct {
}

func NewJobseekerMockRepository() *JobseekerMockRepository {
	return &JobseekerMockRepository{}
}

func (j JobseekerMockRepository) JobSeekers() ([]entity.Jobseeker, error) {
	return []entity.Jobseeker{entity.Jobseekermock1, entity.Jobseekermock2}, nil
}

func (j JobseekerMockRepository) JobSeeker(id int) (entity.Jobseeker, error) {
	if id == 0 {
		return entity.Jobseekermock1, nil
	} else if id == 1 {
		return entity.Jobseekermock2, nil
	}
	return entity.Jobseekermock1, nil
}

func (j JobseekerMockRepository) UpdateJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	return &entity.Jobseekermock2, nil
}

func (j JobseekerMockRepository) DeleteJobSeeker(id int) (entity.Jobseeker, error) {
	return entity.Jobseekermock1, nil
}

func (j JobseekerMockRepository) StoreJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	return &entity.Jobseekermock1, nil
}

func (j JobseekerMockRepository) JsCategories(id int) ([]entity.Category, error) {
	return []entity.Category{entity.Categorymock1, entity.Categorymock2}, nil
}

func (j JobseekerMockRepository) RemoveIntCategory(jsid, jcid int) error {
	panic("implement me")
}

func (j JobseekerMockRepository) AddIntCategory(jsid, jcid int) error {
	return nil
}

func (j JobseekerMockRepository) SetAddress(jsid, addid int) error {
	return nil
}

func (j JobseekerMockRepository) JobseekerByEmail(email string) (entity.Jobseeker, error) {
	if email == "user1.name@gmail.com" {
		return entity.Jobseekermock1, nil
	} else if email == "user2.name@gmail.com" {
		return entity.Jobseekermock2, nil
	}
	return entity.Jobseeker{}, errors.New("error occurred")
}

func (j JobseekerMockRepository) JobseekerByUsername(uname string) (entity.Jobseeker, error) {
	if uname == "user1" {
		return entity.Jobseekermock1, nil
	} else if uname == "user2" {
		return entity.Jobseekermock2, nil
	}
	return entity.Jobseeker{}, errors.New("error occurred")
}

func (j JobseekerMockRepository) ApplicationJobseeker(id int) (entity.Jobseeker, error) {
	if id == 1 {
		return entity.Jobseekermock1, nil
	} else if id == 2 {
		return entity.Jobseekermock2, nil
	}
	return entity.Jobseeker{}, errors.New("error occured")
}

func (j JobseekerMockRepository) UserRoles(user *entity.Jobseeker) ([]entity.Role, []error) {
	return []entity.Role{entity.Rolemock1}, nil
}

func (j JobseekerMockRepository) PhoneExists(phone string) bool {
	return false
}

func (j JobseekerMockRepository) EmailExists(email string) bool {
	return false
}

func (j JobseekerMockRepository) UsernameExists(email string) bool {
	return false
}

func (j JobseekerMockRepository) AlreadyApplied(id uint, id2 uint) bool {
	return false
}
