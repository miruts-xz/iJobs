package application

import "github.com/miruts/iJobs/entity"

type IAppRepository interface {
	Store(*entity.Application) (*entity.Application, error)
	Application(id int) (entity.Application, error)
	UserApplication(jsId int) ([]entity.Application, error)
	ApplicationsOnJob(jobId int) ([]entity.Application, error)
	DeleteApplication(id int) (entity.Application, error)
}
