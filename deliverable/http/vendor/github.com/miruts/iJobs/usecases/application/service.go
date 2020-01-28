package application

import (
	"github.com/miruts/iJobs/entity"
)

// IAppSerivce represents all Application services
type IAppService interface {
	Store(app *entity.Application) (*entity.Application, error)
	Application(id int) (entity.Application, error)
	UserApplication(jsId int) ([]entity.Application, error)
	ApplicationsOnJob(jobId int) ([]entity.Application, error)
	DeleteApplication(id int) (entity.Application, error)
	ApplicationForCompany(cmid int) ([]entity.Application, error)
	UpdateApplication(e *entity.Application) (*entity.Application, error)
}
