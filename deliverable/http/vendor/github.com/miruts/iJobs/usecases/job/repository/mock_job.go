package repository

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
)

type JobMockRepository struct {
}

func NewJobMockRepository() *JobMockRepository {
	return &JobMockRepository{}
}

func (j *JobMockRepository) Jobs() ([]entity.Job, error) {
	panic("implement me")
}

func (j *JobMockRepository) JobsOfCategory(ctgSrv job.CategoryService, catid int) ([]entity.Job, error) {
	return []entity.Job{entity.Jobmock1, entity.Jobmock2}, nil
}

func (j *JobMockRepository) Job(id int) (entity.Job, error) {
	panic("implement me")
}

func (j *JobMockRepository) UpdateJob(job *entity.Job) (*entity.Job, error) {
	panic("implement me")
}

func (j *JobMockRepository) DeleteJob(id int) (entity.Job, error) {
	panic("implement me")
}

func (j *JobMockRepository) StoreJob(job *entity.Job) (*entity.Job, error) {
	panic("implement me")
}

func (j *JobMockRepository) CompanyJobs(cmpSrv company.CompanyService, cmid int) ([]entity.Job, error) {
	panic("implement me")
}
