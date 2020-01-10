package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

type AppGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewApplicationGormRepositoryImpl(conn *gorm.DB) *AppGormRepositoryImpl {
	return &AppGormRepositoryImpl{conn: conn}
}

func (agr *AppGormRepositoryImpl) Store(entity.Application) error {
	return nil
}
func (agr *AppGormRepositoryImpl) UserApplication(jsId int) ([]entity.Application, error) {
	return nil, nil
}
func (agr *AppGormRepositoryImpl) ApplicationsOnJob(jobId int) ([]entity.Application, error) {
	return nil, nil
}
func (agr *AppGormRepositoryImpl) DeleteApplication(id int) error {
	return nil
}
