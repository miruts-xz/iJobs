package service

import (
	"fmt"
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
func (jss *JobseekerServiceImpl) JobSeekers() ([]entity.JobSeeker, error) {
	return jss.jsRepo.JobSeekers()
}

// JobSeeker return jobseeker with a given id
func (jss *JobseekerServiceImpl) JobSeeker(id int) (entity.JobSeeker, error) {
	return jss.jsRepo.JobSeeker(id)
}

// UpdateJobSeeker updates a given jobseeker
func (jss *JobseekerServiceImpl) UpdateJobSeeker(js *entity.JobSeeker) (*entity.JobSeeker, error) {
	return jss.jsRepo.UpdateJobSeeker(js)
}

// DeleteJobSeeker deletes jobseeker with a given id
func (jss *JobseekerServiceImpl) DeleteJobSeeker(id int) (entity.JobSeeker, error) {
	return jss.jsRepo.DeleteJobSeeker(id)
}

// StoreJobSeeker stores new jobseeker
func (jss *JobseekerServiceImpl) StoreJobSeeker(js *entity.JobSeeker) (*entity.JobSeeker, error) {
	return jss.jsRepo.StoreJobSeeker(js)
}

// Suggestions return all jobs related to interested categories given the jobseeker id
func (jss *JobseekerServiceImpl) Suggestions(id int) ([]entity.Job, error) {
	ctgs, err := jss.jsRepo.JsCategories(id)
	var alljobs []entity.Job
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	for _, ctg := range ctgs {
		categjobs, err := jss.jobService.JobsOfCategory(int(ctg.ID))
		if err != nil {
			fmt.Printf("Error: %v", err)
			return alljobs, err
		}
		for _, ctg := range categjobs {
			alljobs = append(alljobs, ctg)
		}
	}
	return alljobs, nil
}
func (jss *JobseekerServiceImpl) SetAddress(jsid, addid int) error {
	return nil
}
func (jss *JobseekerServiceImpl) AddIntCategory(jsid, jcid int) error {
	return nil
}
func (jss *JobseekerServiceImpl) RemoveIntCategory(jsid, jcid int) error {
	return nil
}
