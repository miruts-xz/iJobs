package repository

import "github.com/miruts/iJobs/entity"

type CategoryMockRepository struct {
}

func NewCategoryMockRepository() *CategoryMockRepository {
	return &CategoryMockRepository{}
}

func (c CategoryMockRepository) Categories() ([]entity.Category, error) {
	return []entity.Category{entity.Categorymock1, entity.Categorymock2}, nil
}

func (c CategoryMockRepository) Category(id int) (entity.Category, error) {
	panic("implement me")
}

func (c CategoryMockRepository) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	panic("implement me")
}

func (c CategoryMockRepository) DeleteCategory(id int) (entity.Category, error) {
	panic("implement me")
}

func (c CategoryMockRepository) StoreCategory(category *entity.Category) (*entity.Category, error) {
	panic("implement me")
}
