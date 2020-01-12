package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
)

type JobGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewJobGormRepositoryImpl(conn *gorm.DB) *JobGormRepositoryImpl {
	return &JobGormRepositoryImpl{conn: conn}
}
func (jgr *JobGormRepositoryImpl) Jobs() ([]entity.Job, error) {
	var jobs []entity.Job
	errs := jgr.conn.Find(&jobs).GetErrors()
	if len(errs) > 0 {
		return jobs, errs[0]
	}
	return jobs, nil
}
func (jgr *JobGormRepositoryImpl) JobsOfCategory(ctgSrv job.CategoryService, cat_id int) ([]entity.Job, error) {
	category, err := ctgSrv.Category(cat_id)
	var jobs []entity.Job
	if err != nil {
		fmt.Println(err)
		return jobs, err
	}
	errs := jgr.conn.Model(&category).Related(&jobs, "Jobs").GetErrors()
	if len(errs) > 0 {
		return jobs, errs[0]
	}
	return jobs, nil

}
func (jgr *JobGormRepositoryImpl) Job(id int) (entity.Job, error) {
	var job entity.Job
	errs := jgr.conn.First(&job, id).GetErrors()
	if len(errs) > 0 {
		return job, errs[0]
	}
	return job, nil
}
func (jgr *JobGormRepositoryImpl) UpdateJob(job *entity.Job) (*entity.Job, error) {
	j := job
	errs := jgr.conn.Save(&j).GetErrors()
	if len(errs) > 0 {
		return j, errs[0]
	}
	return j, nil
}
func (jgr *JobGormRepositoryImpl) DeleteJob(id int) (entity.Job, error) {
	job, err := jgr.Job(id)
	if err != nil {
		fmt.Println(err)
		return job, err
	}
	errs := jgr.conn.Delete(&job).GetErrors()
	if len(errs) > 0 {
		return job, errs[0]
	}
	return job, nil
}
func (jgr *JobGormRepositoryImpl) StoreJob(job *entity.Job) (*entity.Job, error) {
	j := job
	errs := jgr.conn.Create(&j).GetErrors()
	if len(errs) > 0 {
		return j, errs[0]
	}
	return j, nil
}
func (jgr *JobGormRepositoryImpl) CompanyJobs(cmpSrv company.CompanyService, cmid int) ([]entity.Job, error) {
	company, err := cmpSrv.Company(cmid)
	var jobs []entity.Job
	if err != nil {
		fmt.Println(err)
		return jobs, err
	}
	errs := jgr.conn.Model(&company).Related(&jobs, "Jobs").GetErrors()
	if len(errs) > 0 {
		return jobs, errs[0]
	}
	return jobs, nil
}
