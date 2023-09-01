package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type ProductUseCase interface {
	CreateNew(payload model.Product) error
	FindById(id string) (model.Product, error)
	FindAll() ([]model.Product, error)
	GetByName(name string) ([]model.Product, error)
	Update(payload model.Product) error
	Delete(id string) error
}

type productUseCase struct {
	repo       repository.ProductRepository
	uomUseCase UomUseCase
}

// CreateNew implements ProductUseCase.
func (p *productUseCase) CreateNew(payload model.Product) error {
	// Handler id tidak duplicat dilakukan pada db constraint uniqe
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if payload.Uom.Id == "" {
		return fmt.Errorf("uom_id is required")
	}

	err := p.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new Products %v", err)
	}
	fmt.Println("New product has been saved..")
	return nil
}

// DeleteById implements ProductUseCase.
func (p *productUseCase) Delete(id string) error {
	err := p.repo.DeleteById(id)
	if err != nil {
		return err
	}
	fmt.Printf("Product with id: %s has been deleted\n", id)
	return nil
}

// FindAll implements ProductUseCase.
func (p *productUseCase) FindAll() ([]model.Product, error) {
	return p.repo.FindAll()
}

// FindById implements ProductUseCase.
func (p *productUseCase) FindById(id string) (model.Product, error) {
	product, err := p.repo.FindById(id)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// GetByName implements ProductUseCase.
func (p *productUseCase) GetByName(name string) ([]model.Product, error) {
	return p.repo.FindByName(name)
}

// Update implements ProductUseCase.
func (p *productUseCase) Update(payload model.Product) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price <= 0 {
		return fmt.Errorf("price must be greater than zero")
	}

	_, err := p.uomUseCase.FindById(payload.Uom.Id)
	if err != nil {
		return err
	}

	_, err = p.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = p.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}

	return nil
}

func NewProductUseCase(repo repository.ProductRepository, uom UomUseCase) ProductUseCase {
	return &productUseCase{
		repo:       repo,
		uomUseCase: uom,
	}
}
