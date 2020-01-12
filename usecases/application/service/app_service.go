package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/application/repository"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

type AppService struct {
	appRepo *repository.AppGormRepositoryImpl
	jsSrv   jobseeker.JobseekerService
	jobSrv  job.JobService
	cmpSrv  company.CompanyService
}

func NewAppService(appRepo *repository.AppGormRepositoryImpl, jsSrv jobseeker.JobseekerService, jobSrv job.JobService, cmpSrv company.CompanyService) *AppService {
	return &AppService{appRepo: appRepo, jsSrv: jsSrv, jobSrv: jobSrv, cmpSrv: cmpSrv}
}

func (appService *AppService) Application(id int) (entity.Application, error) {
	return appService.appRepo.Application(id)
}
func (appService *AppService) Store(app *entity.Application) error {
	return appService.appRepo.Store(app)
}
func (appService *AppService) UserApplication(JsId int) ([]entity.Application, error) {
	return appService.appRepo.UserApplication(appService.jsSrv, JsId)
}
func (appService *AppService) ApplicationsOnJob(jobId int) ([]entity.Application, error) {
	return appService.appRepo.ApplicationsOnJob(appService.jobSrv, jobId)
}
func (appService *AppService) DeleteApplication(id int) (entity.Application, error) {
	return appService.appRepo.DeleteApplication(id)
}
func (appService *AppService) ApplicationForCompany(cmid int) ([]entity.Application, error) {
	return appService.appRepo.ApplicationForCompany(cmid)
}
