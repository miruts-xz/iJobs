package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

// CategoryGormRepositoryImpl implements the menu.CategoryRepository interface
type CategoryGormRepositoryImpl struct {
	conn *gorm.DB
}

// NewCategoryGormRepositoryImpl will create an object of CategoryGormRepositoryImpl
func NewCategoryGormRepositoryImpl(conn *gorm.DB) *CategoryGormRepositoryImpl {
	return &CategoryGormRepositoryImpl{conn: conn}
}

// Categories returns all cateogories from the database
func (cgr *CategoryGormRepositoryImpl) Categories() ([]entity.Category, error) {
	var categories []entity.Category
	errs := cgr.conn.Find(&categories).GetErrors()
	if len(errs) > 0 {
		return categories, nil
	}
	return categories, nil
}

//Category returns a category with a given id
func (cgr *CategoryGormRepositoryImpl) Category(id int) (entity.Category, error) {
	var category entity.Category
	errs := cgr.conn.First(&category, id).GetErrors()
	if len(errs) > 0 {
		return category, errs[0]
	}
	return category, nil
}

//UpdateCategory updates a given object with a new data
func (cgr *CategoryGormRepositoryImpl) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	cat := category
	errs := cgr.conn.Save(&cat).GetErrors()
	if len(errs) > 0 {
		return cat, errs[0]
	}
	return cat, nil
}

//DeleteCategory removes a category from a database by its id
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

//StoreCategory stores new category information to database
func (cgr *CategoryGormRepositoryImpl) StoreCategory(category *entity.Category) (*entity.Category, error) {
	cat := category
	errs := cgr.conn.Create(cat).GetErrors()
	if len(errs) > 0 {
		return cat, errs[0]
	}
	return cat, nil
}
