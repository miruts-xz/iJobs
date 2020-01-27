package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
)

// AppGormRepositoryImpl represents IAppRepository interface
type AppGormRepositoryImpl struct {
	conn *gorm.DB
}

// NewAppGormRepositoryImpl creates new AppGormRepositoryImpl
func NewAppGormRepositoryImpl(conn *gorm.DB) *AppGormRepositoryImpl {
	return &AppGormRepositoryImpl{conn: conn}
}

// Store stores application
func (agr *AppGormRepositoryImpl) Store(app *entity.Application) (*entity.Application, error) {
	application := app
	errs := agr.conn.Create(application).GetErrors()
	if len(errs) > 0 {
		return application, errs[0]
	}
	return application, nil
}

// Application finds application by id
func (agr *AppGormRepositoryImpl) Application(id int) (entity.Application, error) {
	var application entity.Application
	errs := agr.conn.First(&application, id).GetErrors()
	if len(errs) > 0 {
		return application, errs[0]
	}
	return application, nil
}

// UserApplication finds all application given jobseeker id and service
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

// ApplicationsOnJob retrieves all Application on a given job
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

// Delete Application deletes application with given id
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

// ApplicationForCompany retrieves all job-applications for a given company
func (agr *AppGormRepositoryImpl) ApplicationForCompany(cmid int) ([]entity.Application, error) {
	var jobs []entity.Job
	errs := agr.conn.Where("company_id = ?", cmid).Find(&jobs).GetErrors()
	var tobereturned []entity.Application
	var applications []entity.Application

	for i, _ := range jobs {
		errs := agr.conn.Where("job_id = ?", jobs[i].ID).Find(&applications).GetErrors()
		if len(errs) > 0 {
			fmt.Println(errs)
			return tobereturned, errs[0]
		}
		tobereturned = append(tobereturned, applications...)
	}
	if len(errs) > 0 {
		return tobereturned, errs[0]
	}
	return tobereturned, nil
}
