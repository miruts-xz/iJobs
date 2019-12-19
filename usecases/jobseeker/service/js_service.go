package service

import (
	"fmt"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

type JobseekerServiceImpl struct {
	jsRepo     jobseeker.JobseekerRepository
	jobService job.JobService
}

func (jss *JobseekerServiceImpl) NewJobseekerServiceImpl(jsr jobseeker.JobseekerRepository, jobs job.JobService) *JobseekerServiceImpl {
	return &JobseekerServiceImpl{jsRepo: jsr, jobService: jobs}
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
