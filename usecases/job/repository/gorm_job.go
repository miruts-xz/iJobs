package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

type JobGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewJobGormRepositoryImpl(conn *gorm.DB) *JobGormRepositoryImpl {
	return &JobGormRepositoryImpl{conn: conn}
}
func (jgr *JobGormRepositoryImpl) Jobs() ([]entity.Job, error) {
	return nil, nil
}
func (jgr *JobGormRepositoryImpl) JobsOfCategory(cat_id int) ([]entity.Job, error) {
	return nil, nil
}
func (jgr *JobGormRepositoryImpl) Job(id int) (entity.Job, error) {
	return entity.Job{}, nil
}
func (jgr *JobGormRepositoryImpl) UpdateJob(job entity.Job) error {
	return nil
}
func (jgr *JobGormRepositoryImpl) DeleteJob(id int) error {
	return nil
}
func (jgr *JobGormRepositoryImpl) StoreJob(job entity.Job) error {
	return nil
}
