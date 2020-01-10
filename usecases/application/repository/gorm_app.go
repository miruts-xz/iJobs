package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

type AppGormRepositoryImpl struct {
	conn   *gorm.DB
	jsSrv  jobseeker.JobseekerService
	jobSrv job.JobService
}

func NewAppGormRepositoryImpl(conn *gorm.DB, jsSrv jobseeker.JobseekerService, jobSrv job.JobService) *AppGormRepositoryImpl {
	return &AppGormRepositoryImpl{conn: conn, jsSrv: jsSrv, jobSrv: jobSrv}
}

func (agr *AppGormRepositoryImpl) Store(app *entity.Application) (*entity.Application, error) {
	application := app
	errs := agr.conn.Create(application).GetErrors()
	if len(errs) > 0 {
		return application, errs[0]
	}
	return application, nil
}
func (agr *AppGormRepositoryImpl) Application(id int) (entity.Application, error) {
	var application entity.Application
	errs := agr.conn.First(&application, id).GetErrors()
	if len(errs) > 0 {
		return application, errs[0]
	}
	return application, nil
}
func (agr *AppGormRepositoryImpl) UserApplication(jsId int) ([]entity.Application, error) {
	jobseeker, err := agr.jsSrv.JobSeeker(jsId)
	var applications []entity.Application
	if err != nil {
		fmt.Println(err)
		return applications, err
	}
	errs := agr.conn.Model(&jobseeker).Related(&applications, "Applications").GetErrors()
	if len(errs) > 0 {
		return applications, errs[0]
	}
	return applications, nil
}
func (agr *AppGormRepositoryImpl) ApplicationsOnJob(jobId int) ([]entity.Application, error) {
	job, err := agr.jobSrv.Job(jobId)
	var applications []entity.Application
	if err != nil {
		fmt.Println(err)
		return applications, err
	}
	errs := agr.conn.Model(job).Related(&applications, "Applications").GetErrors()
	if len(errs) > 0 {
		return applications, errs[0]
	}
	return applications, nil
}
func (agr *AppGormRepositoryImpl) DeleteApplication(id int) (entity.Application, error) {
	application, err := agr.Application(id)
	if err != nil {
		return application, err
	}
	errs := agr.conn.Delete(application, id).GetErrors()
	if len(errs) > 0 {
		return application, errs[0]
	}
	return application, nil
}
