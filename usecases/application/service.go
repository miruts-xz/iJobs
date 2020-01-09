package application

import "github.com/miruts/iJobs/entity"

type IAppService interface {
	Store(entity.Application) error
	UserApplication(JsId int) ([]entity.Application, error)
	ApplicationsOnJob(jobId int) ([]entity.Application, error)
	DeleteApplication(id int) error
}
