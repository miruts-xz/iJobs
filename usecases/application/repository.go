package application

import "github.com/miruts/iJobs/entity"

type IAppRepository interface {
	Store(entity.Application) error
	Application(appId int) ([]entity.Application, error)
	UserApplication(jsId int) ([]entity.Application, error)
	ApplicationsOnJob(jobId int) ([]entity.Application, error)
	DeleteApplication(id int) error
}
