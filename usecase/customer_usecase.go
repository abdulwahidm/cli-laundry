package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type CustomerUseCase interface {
	CreateNew(payload model.Customer) error
	FindById(id string) (model.Customer, error)
	FindAll() ([]model.Customer, error)
	GetByName(name string) ([]model.Customer, error)
	Update(payload model.Customer) error
	DeleteById(id string) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

// CreateNew implements CustomerUseCase.
func (c *customerUseCase) CreateNew(payload model.Customer) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := c.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("faield to create new customers: %v", err)
	}
	fmt.Println("New customer has been added")
	return nil

}

// DeleteById implements CustomerUseCase.
func (c *customerUseCase) DeleteById(id string) error {
	panic("unimplemented")
}

// FindAll implements CustomerUseCase.
func (c *customerUseCase) FindAll() ([]model.Customer, error) {
	return c.repo.FindAll()
}

// FindById implements CustomerUseCase.
func (c *customerUseCase) FindById(id string) (model.Customer, error) {
	panic("unimplemented")
}

// GetByName implements CustomerUseCase.
func (c *customerUseCase) GetByName(name string) ([]model.Customer, error) {
	panic("unimplemented")
}

// Update implements CustomerUseCase.
func (c *customerUseCase) Update(payload model.Customer) error {
	panic("unimplemented")
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
