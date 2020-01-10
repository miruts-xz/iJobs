package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

type CategoryGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewCategoryGormRepositoryImpl(conn *gorm.DB) *CategoryGormRepositoryImpl {
	return &CategoryGormRepositoryImpl{conn: conn}
}
func (cgr *CategoryGormRepositoryImpl) Categories() ([]entity.Category, error) {
	return nil, nil
}
func (cgr *CategoryGormRepositoryImpl) Category(id int) (entity.Category, error) {
	return entity.Category{}, nil
}
func (cgr *CategoryGormRepositoryImpl) UpdateCategory(category entity.Category) error {
	return nil
}
func (cgr *CategoryGormRepositoryImpl) DeleteCategory(id int) error {
	return nil
}
func (cgr *CategoryGormRepositoryImpl) StoreCategory(category entity.Category) error {
	return nil
}
