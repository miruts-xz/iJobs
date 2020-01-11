package repository

import (
	"fmt"
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
	var categories []entity.Category
	errs := cgr.conn.Find(&categories).GetErrors()
	if len(errs) > 0 {
		return categories, nil
	}
	return categories, nil
}
func (cgr *CategoryGormRepositoryImpl) Category(id int) (entity.Category, error) {
	var category entity.Category
	errs := cgr.conn.Find(&category, id).GetErrors()
	if len(errs) > 0 {
		return category, errs[0]
	}
	return category, nil
}
func (cgr *CategoryGormRepositoryImpl) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	cat := category
	errs := cgr.conn.Save(&cat).GetErrors()
	if len(errs) > 0 {
		return cat, errs[0]
	}
	return cat, nil
}
func (cgr *CategoryGormRepositoryImpl) DeleteCategory(id int) (entity.Category, error) {
	category, err := cgr.Category(id)
	if err != nil {
		fmt.Println(err)
		return category, err
	}
	errs := cgr.conn.Delete(&category).GetErrors()
	if len(errs) > 0 {
		return category, errs[0]
	}
	return category, nil
}
func (cgr *CategoryGormRepositoryImpl) StoreCategory(category *entity.Category) (*entity.Category, error) {
	cat := category
	errs := cgr.conn.Create(cat).GetErrors()
	if len(errs) > 0 {
		return cat, errs[0]
	}
	return cat, nil
}
