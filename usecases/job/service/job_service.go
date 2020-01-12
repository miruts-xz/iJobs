package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/job/repository"
)

// JobServices Implements JobService Interface
type JobServices struct {
	handler *repository.JobGormRepositoryImpl
	ctgSrv  job.CategoryService
}

// NewJobServices creates new JobServices
func NewJobServices(handler *repository.JobGormRepositoryImpl, ctgSrv job.CategoryService) *JobServices {
	return &JobServices{handler: handler, ctgSrv: ctgSrv}
}

// Jobs retrieves all jobs
func (jobService *JobServices) Jobs() ([]entity.Job, error) {
	return jobService.handler.Jobs()

}

//Returns all jobs under a specific category
func (jobService *JobServices) JobsOfCategory(cat_id int) ([]entity.Job, error) {
	return jobService.handler.JobsOfCategory(jobService.ctgSrv, cat_id)
}

//Returns a job given an its id
func (jobService *JobServices) Job(id int) (entity.Job, error) {
	return jobService.handler.Job(id)

}

//Updates a job given the udpated job object
func (jobService *JobServices) UpdateJob(job *entity.Job) (*entity.Job, error) {
	return jobService.handler.UpdateJob(job)
}

//Deletes a job given its id
func (jobService *JobServices) DeleteJob(id int) (entity.Job, error) {
	return jobService.handler.DeleteJob(id)

}

//Adds a job to the database
func (jobService *JobServices) StoreJob(job *entity.Job) (*entity.Job, error) {
	return jobService.handler.StoreJob(job)

}

// CompanyJobs retrieves jobs of specific company
func (jobService *JobServices) CompanyJobs(cmpSrv company.CompanyService, cmid int) ([]entity.Job, error) {
	return jobService.handler.CompanyJobs(cmpSrv, cmid)
}
