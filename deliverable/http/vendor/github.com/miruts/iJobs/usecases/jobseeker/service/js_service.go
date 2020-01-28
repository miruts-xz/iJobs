package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

// JobseekerServiceImpl implements JobseekerService interface
type JobseekerServiceImpl struct {
	jsRepo     jobseeker.JobseekerRepository
	jobService job.JobService
}

// NewJobseekerServiceImpl returns new JobseekerServiceImpl
func NewJobseekerServiceImpl(jsr jobseeker.JobseekerRepository, jobs job.JobService) *JobseekerServiceImpl {
	return &JobseekerServiceImpl{jsRepo: jsr, jobService: jobs}
}

// JobSeekers return all jobseekers
func (jss *JobseekerServiceImpl) JobSeekers() ([]entity.Jobseeker, error) {
	return jss.jsRepo.JobSeekers()
}

// JobSeeker return jobseeker with a given id
func (jss *JobseekerServiceImpl) JobSeeker(id int) (entity.Jobseeker, error) {
	return jss.jsRepo.JobSeeker(id)
}

// UpdateJobSeeker updates a given jobseeker
func (jss *JobseekerServiceImpl) UpdateJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	return jss.jsRepo.UpdateJobSeeker(js)
}

// DeleteJobSeeker deletes jobseeker with a given id
func (jss *JobseekerServiceImpl) DeleteJobSeeker(id int) (entity.Jobseeker, error) {
	return jss.jsRepo.DeleteJobSeeker(id)
}

// StoreJobSeeker stores new jobseeker
func (jss *JobseekerServiceImpl) StoreJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	return jss.jsRepo.StoreJobSeeker(js)
}

// Suggestions return all jobs related to interested categories given the jobseeker id
func (jss *JobseekerServiceImpl) Suggestions(id int) ([]entity.Job, error) {
	ctgs, err := jss.jsRepo.JsCategories(id)
	var alljobs []entity.Job
	if err != nil {
		return nil, err
	}
	for _, ctg := range ctgs {
		categjobs, err := jss.jobService.JobsOfCategory(int(ctg.ID))
		if err != nil {
			return alljobs, err
		}
		for _, ctg := range categjobs {
			alljobs = append(alljobs, ctg)
		}
	}
	return alljobs, nil
}

// SetAddress sets address of jobseeker
func (jss *JobseekerServiceImpl) SetAddress(jsid, addid int) error {
	return jss.jsRepo.SetAddress(jsid, addid)
}

// AddIntCategory Adds interested job category
func (jss *JobseekerServiceImpl) AddIntCategory(jsid, jcid int) error {
	return jss.jsRepo.AddIntCategory(jsid, jcid)
}

// RemoveIntCategory removes category from interested list
func (jss *JobseekerServiceImpl) RemoveIntCategory(jsid, jcid int) error {
	return jss.RemoveIntCategory(jsid, jcid)
}

// JobseekerByEmail retrieves a jobseeker given email
func (jss *JobseekerServiceImpl) JobseekerByEmail(email string) (entity.Jobseeker, error) {
	return jss.jsRepo.JobseekerByEmail(email)
}

// JobseekerByUsername retrieves a jobseeker given username
func (jss *JobseekerServiceImpl) JobseekerByUsername(uname string) (entity.Jobseeker, error) {
	return jss.jsRepo.JobseekerByUsername(uname)
}

// ApplicationJobseekers retrieves jobseekers who applied for application of given id
func (jss *JobseekerServiceImpl) ApplicationJobseeker(id int) (entity.Jobseeker, error) {
	return jss.jsRepo.ApplicationJobseeker(id)
}

// UserRoles returns list of roles a user has
func (jss *JobseekerServiceImpl) UserRoles(user *entity.Jobseeker) ([]entity.Role, []error) {
	userRoles, errs := jss.jsRepo.UserRoles(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}

// PhoneExists check if there is a user with a given phone number
func (jss *JobseekerServiceImpl) PhoneExists(phone string) bool {
	exists := jss.jsRepo.PhoneExists(phone)
	return exists
}
func (jss *JobseekerServiceImpl) UsernameExists(email string) bool {
	exists := jss.jsRepo.UsernameExists(email)
	return exists
}

// EmailExists checks if there exist a user with a given email address
func (jss *JobseekerServiceImpl) EmailExists(email string) bool {
	exists := jss.jsRepo.EmailExists(email)
	return exists
}
func (jss *JobseekerServiceImpl) AlreadyApplied(id uint, id2 uint) bool {
	applied := jss.jsRepo.AlreadyApplied(id, id2)
	return applied
}
