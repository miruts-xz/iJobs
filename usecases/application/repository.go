package application

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

// IAppRepository represents Application data repository
type IAppRepository interface {
	Store(app *entity.Application) error
	Application(id int) (entity.Application, error)
	UserApplication(jsSrv jobseeker.JobseekerService, jsId int) ([]entity.Application, error)
	ApplicationsOnJob(jobSrv job.JobService, jobId int) ([]entity.Application, error)
	DeleteApplication(id int) (entity.Application, error)
	ApplicationForCompany(cmid int) ([]entity.Application, error)
}
