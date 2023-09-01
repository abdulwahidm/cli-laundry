package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type EmployeeUseCase interface {
	CreateNew(payload model.Employee) error
	FindById(id string) (model.Employee, error)
	FindAll() ([]model.Employee, error)
	GetByName(name string) ([]model.Employee, error)
	Update(payload model.Employee) error
	Delete(id string) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

// CreateNew implements EmployeeUseCase.
func (e *employeeUseCase) CreateNew(payload model.Employee) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := e.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("faield to create new employee: %v", err)
	}
	fmt.Println("New employee has been added")
	return nil
}

// Delete implements EmployeeUseCase.
func (e *employeeUseCase) Delete(id string) error {
	_, err := e.FindById(id)
	if err != nil {
		return err
	}

	err = e.repo.DeleteById(id)
	if err != nil {
		return err
	}
	fmt.Printf("Employee with id: %s deleted\n", id)
	return nil
}

// FindAll implements EmployeeUseCase.
func (e *employeeUseCase) FindAll() ([]model.Employee, error) {
	return e.repo.FindAll()
}

// FindById implements EmployeeUseCase.
func (e *employeeUseCase) FindById(id string) (model.Employee, error) {
	employee, err := e.repo.FindById(id)
	if err != nil {
		return model.Employee{}, fmt.Errorf("employee with id: %s not found.", id)
	}
	return employee, nil
}

// GetByName implements EmployeeUseCase.
func (e *employeeUseCase) GetByName(name string) ([]model.Employee, error) {
	return e.repo.FindByName(name)
}

// Update implements EmployeesUseCase.
func (e *employeeUseCase) Update(payload model.Employee) error {
	_, err := e.FindById(payload.Id)
	if err != nil {
		return err
	}
	if payload.Id == "" || payload.Name == "" || payload.PhoneNumber == "" {
		return fmt.Errorf("data is required")
	}

	err = e.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom: %v", err)
	}
	return nil
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
