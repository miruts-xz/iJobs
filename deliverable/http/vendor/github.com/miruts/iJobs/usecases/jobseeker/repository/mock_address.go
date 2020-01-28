package repository

import "github.com/miruts/iJobs/entity"

type AddressMockRepository struct {
}

func NewAddressMockRepository() *AddressMockRepository {
	return &AddressMockRepository{}
}

func (a AddressMockRepository) Addresses() ([]entity.Address, error) {
	panic("implement me")
}

func (a AddressMockRepository) Address(id int) (entity.Address, error) {
	panic("implement me")
}

func (a AddressMockRepository) UpdateAddress(category *entity.Address) (*entity.Address, error) {
	panic("implement me")
}

func (a AddressMockRepository) DeleteAddress(id int) (entity.Address, error) {
	panic("implement me")
}

func (a AddressMockRepository) StoreAddress(category *entity.Address) (*entity.Address, error) {
	panic("implement me")
}
