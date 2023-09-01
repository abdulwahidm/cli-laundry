package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type UomUseCase interface {
	CreateNew(payload model.Uom) error
	FindById(id string) (model.Uom, error)
	FindAll() ([]model.Uom, error)
	Update(payload model.Uom) error
	Delete(id string) error
}

type uomUseCase struct {
	repo repository.UomRepository
}

// CreateNew implements UomUseCase.
func (u *uomUseCase) CreateNew(payload model.Uom) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	fmt.Println("New uom has been saved")
	return u.repo.Save(payload)
}

// Delete implements UomUseCase.
func (u *uomUseCase) Delete(id string) error {
	uom, err := u.FindById(id)
	if err != nil {
		return err
	}

	err = u.repo.DeleteById(uom.Id)
	if err != nil {
		return fmt.Errorf("failed to delete uom: %v", err)
	}
	fmt.Printf("Uom with id %s has been deleted\n", id)
	return nil
}

// FindAll implements UomUseCase.
func (u *uomUseCase) FindAll() ([]model.Uom, error) {
	return u.repo.FindAll()
}

// FindById implements UomUseCase.
func (u *uomUseCase) FindById(id string) (model.Uom, error) {
	uom, err := u.repo.FindById(id)
	if err != nil {
		return model.Uom{}, fmt.Errorf("uom with id: %s not found", id)
	}
	return uom, nil
}

// Update implements UomUseCase.
func (u *uomUseCase) Update(payload model.Uom) error {
	_, err := u.FindById(payload.Id)
	if err != nil {
		return err
	}

	if payload.Id == "" || payload.Name == "" {
		return fmt.Errorf("data is required")
	}
	return u.repo.Update(payload)
}

func NewUomUseCase(repo repository.UomRepository) UomUseCase {
	return &uomUseCase{repo: repo}
}
