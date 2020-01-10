package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/application"
)

type AppService struct {
	appRepo application.IAppRepository
}

func NewAppservice(appRepo application.IAppRepository) *AppService {
	return &AppService{appRepo: appRepo}
}

func (appService *AppService) Store(app entity.Application) error {
	//return appService.appRepo.Store(app)
	return nil
}
func (appService *AppService) UserApplication(JsId int) ([]entity.Application, error) {
	//return appService.appRepo.UserApplication(JsId)
	return nil, nil
}
func (appService *AppService) ApplicationsOnJob(jobId int) ([]entity.Application, error) {
	//return appService.appRepo.ApplicationsOnJob(jobId)
	return nil, nil
}
func (appService *AppService) DeleteApplication(id int) error {
	//return appService.DeleteApplication(id)
	return nil
}
