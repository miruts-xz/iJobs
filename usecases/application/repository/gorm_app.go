package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

type AppGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewAppGormRepositoryImpl(conn *gorm.DB) *AppGormRepositoryImpl {
	return &AppGormRepositoryImpl{conn: conn}
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
func (agr *AppGormRepositoryImpl) UserApplication(jsSrv jobseeker.JobseekerService, jsId int) ([]entity.Application, error) {
	jobseeker, err := jsSrv.JobSeeker(jsId)
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
func (agr *AppGormRepositoryImpl) ApplicationsOnJob(jobSrv job.JobService, jobId int) ([]entity.Application, error) {
	job, err := jobSrv.Job(jobId)
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
