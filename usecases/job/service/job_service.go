package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
)

type JobServices struct {
	handler job.JobRepository
}

func NewJobService(handler job.JobRepository) *JobServices {
	return &JobServices{handler: handler}
}

func (jobService *JobServices) Jobs() ([]entity.Job, error) {
	return jobService.handler.Jobs()
	
}

//Returns all jobs under a specific category
func (jobService *JobServices) JobsOfCategory(cat_id int) ([]entity.Job, error) {
	return jobService.handler.JobsOfCategory(cat_id)

}

//Returns a job given an its id
func (jobService *JobServices) Job(id int) (entity.Job, error) {
	return jobService.handler.Job(id)
	
}

//Updates a job given the udpated job object
func (jobService *JobServices) UpdateJob(job entity.Job) error {
	return jobService.UpdateJob(job)
	

//Deletes a job given its id
func (jobService *JobServices) DeleteJob(id int) error {
	return jobService.handler.DeleteJob(id)
	
}

//Adds a job to the database
func (jobService *JobServices) StoreJob(job entity.Job) error {
	return jobService.handler.StoreJob(job)

}
