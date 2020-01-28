package repository

import (
	"github.com/miruts/iJobs/entity"
)

type RoleMockRepository struct {
}

func NewRoleMockRepository() *RoleMockRepository {
	return &RoleMockRepository{}
}

func (r RoleMockRepository) Roles() ([]entity.Role, []error) {
	panic("implement me")
}

func (r RoleMockRepository) Role(id uint) (*entity.Role, []error) {
	panic("implement me")
}

func (r RoleMockRepository) RoleByName(name string) (*entity.Role, []error) {
	panic("implement me")
}

func (r RoleMockRepository) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	panic("implement me")
}

func (r RoleMockRepository) DeleteRole(id uint) (*entity.Role, []error) {
	panic("implement me")
}

func (r RoleMockRepository) StoreRole(role *entity.Role) (*entity.Role, []error) {
	panic("implement me")
}
