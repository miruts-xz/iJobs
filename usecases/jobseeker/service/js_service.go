package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

type JobseekerServiceImpl struct {
	jsRepo jobseeker.JobseekerRepository
}

func (jss *JobseekerServiceImpl) NewJobseekerServiceImpl(jsr jobseeker.JobseekerRepository) *JobseekerServiceImpl {
	return &JobseekerServiceImpl{jsRepo: jsr}
}
func (jss *JobseekerServiceImpl) JobSeekers() ([]entity.JobSeeker, error) {
	return jss.jsRepo.JobSeekers()
}
func (jss *JobseekerServiceImpl) JobSeeker(id int) (entity.JobSeeker, error) {
	return jss.jsRepo.JobSeeker(id)
}
func (jss *JobseekerServiceImpl) UpdateJobSeeker(js entity.JobSeeker) error {
	return jss.jsRepo.UpdateJobSeeker(js)
}
func (jss *JobseekerServiceImpl) DeleteJobSeeker(id int) error {
	return jss.jsRepo.DeleteJobSeeker(id)
}
func (jss *JobseekerServiceImpl) StoreJobSeeker(js entity.JobSeeker) error {
	return jss.jsRepo.StoreJobSeeker(js)
}
func (jss *JobseekerServiceImpl) Suggestions(id int) ([]entity.Job, error) {
	//todo
	return nil, nil
}
