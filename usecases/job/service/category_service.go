package service

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
)

// CategoryServiceImpl implements menu.CategoryService interface
type CategoryServiceImpl struct {
	categoryRepo job.CategoryRepository
}

// NewCategoryServiceImpl will create new CategoryService object
func NewCategoryServiceImpl(CatRepo job.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{categoryRepo: CatRepo}
}

// Categories returns list of categories
func (cs *CategoryServiceImpl) Categories() ([]entity.Category, error) {

	categories, err := cs.categoryRepo.Categories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// StoreCategory persists new category information
func (cs *CategoryServiceImpl) StoreCategory(category *entity.Category) (*entity.Category, error) {

	return cs.categoryRepo.StoreCategory(category)

}

// Category returns a category object with a given id
func (cs *CategoryServiceImpl) Category(id int) (entity.Category, error) {

	c, err := cs.categoryRepo.Category(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateCategory updates a cateogory with new data
func (cs *CategoryServiceImpl) UpdateCategory(category *entity.Category) (*entity.Category, error) {

	return cs.categoryRepo.UpdateCategory(category)

}

// DeleteCategory delete a category by its id
func (cs *CategoryServiceImpl) DeleteCategory(id int) (entity.Category, error) {

	return cs.categoryRepo.DeleteCategory(id)
}
