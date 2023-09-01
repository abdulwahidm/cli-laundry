package repository

import (
	"database/sql"
	"fmt"
	"go-laundry/model"
)

type CustomerRepository interface {
	Save(customer model.Customer) error
	FindAll() ([]model.Customer, error)
	FindById(id string) (model.Customer, error)
	FindByName(name string) (model.Customer, error)
	DeleteById(id string) error
}

type customerRepository struct {
	db *sql.DB
}

// DeleteById implements CustomerRepository.
func (c *customerRepository) DeleteById(id string) error {
	panic("unimplemented")
}

// FindAll implements CustomerRepository.
func (c *customerRepository) FindAll() ([]model.Customer, error) {
	rows, err := c.db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.PhoneNumber,
			&customer.Address)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

// FindById implements CustomerRepository.
func (c *customerRepository) FindById(id string) (model.Customer, error) {
	panic("unimplemented")
}

// FindByName implements CustomerRepository.
func (c *customerRepository) FindByName(name string) (model.Customer, error) {
	panic("unimplemented")
}

// Save implements CustomerRepository.
func (c *customerRepository) Save(customer model.Customer) error {
	_, err := c.db.Exec("INSERT INTO customer VALUES ($1, $2, $3, $4)",
		customer.Id, customer.Name, customer.PhoneNumber, customer.Address)
	if err != nil {
		fmt.Println("error from db")
		return err
	}
	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
