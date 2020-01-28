package repository

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

type ApplicationMockRepository struct {
}

func NewApplicationMockRepository() *ApplicationMockRepository {
	return &ApplicationMockRepository{}
}

func (a ApplicationMockRepository) Store(app *entity.Application) (*entity.Application, error) {
	panic("implement me")
}

func (a ApplicationMockRepository) Application(id int) (entity.Application, error) {
	panic("implement me")
}

func (a ApplicationMockRepository) UserApplication(jsSrv jobseeker.JobseekerService, jsId int) ([]entity.Application, error) {
	return []entity.Application{entity.Applicatiomock1, entity.Applicatiomock2}, nil
}

func (a ApplicationMockRepository) ApplicationsOnJob(jobSrv job.JobService, jobId int) ([]entity.Application, error) {
	panic("implement me")
}

func (a ApplicationMockRepository) DeleteApplication(id int) (entity.Application, error) {
	panic("implement me")
}

func (a ApplicationMockRepository) ApplicationForCompany(cmid int) ([]entity.Application, error) {
	panic("implement me")
}

func (a ApplicationMockRepository) UpdateApplication(e *entity.Application) (*entity.Application, error) {
	panic("implement me")
}
