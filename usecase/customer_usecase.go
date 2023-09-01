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
	Delete(id string) error
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
func (c *customerUseCase) Delete(id string) error {
	_, err := c.FindById(id)
	if err != nil {
		return err
	}

	err = c.repo.DeleteById(id)
	if err != nil {
		return err
	}
	fmt.Printf("Customer with id: %s deleted", id)
	return nil
}

// FindAll implements CustomerUseCase.
func (c *customerUseCase) FindAll() ([]model.Customer, error) {
	return c.repo.FindAll()
}

// FindById implements CustomerUseCase.
func (c *customerUseCase) FindById(id string) (model.Customer, error) {
	customer, err := c.repo.FindById(id)
	if err != nil {
		return model.Customer{}, fmt.Errorf("customer with id: %s not found.", id)
	}
	return customer, nil
}

// GetByName implements CustomerUseCase.
func (c *customerUseCase) GetByName(name string) ([]model.Customer, error) {
	panic("unimplemented")
}

// Update implements CustomerUseCase.
func (c *customerUseCase) Update(payload model.Customer) error {
	_, err := c.FindById(payload.Id)
	if err != nil {
		return err
	}
	if payload.Id == "" || payload.Name == "" || payload.PhoneNumber == "" {
		return fmt.Errorf("data is required")
	}

	err = c.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom: %v", err)
	}
	return nil
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
