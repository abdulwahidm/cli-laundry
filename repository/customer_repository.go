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
	Update(customer model.Customer) error
	DeleteById(id string) error
}

type customerRepository struct {
	db *sql.DB
}

// Update implements CustomerRepository.
func (c *customerRepository) Update(customer model.Customer) error {
	_, err := c.db.Exec("UPDATE customer SET name = $2, phone_number = $3,address = $4 WHERE id = $1",
		customer.Id,
		customer.Name,
		customer.PhoneNumber,
		customer.Address,
	)
	if err != nil {
		return err
	}
	fmt.Println("update customer")
	return nil
}

// DeleteById implements CustomerRepository.
func (c *customerRepository) DeleteById(id string) error {
	_, err := c.db.Exec("DELETE FROM customer WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
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
	row := c.db.QueryRow("SELECT * FROM customer WHERE id = $1", id)
	var customer model.Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

// FindByName implements CustomerRepository.
func (c *customerRepository) FindByName(name string) (model.Customer, error) {
	panic("unimplemented")
}

// Save implements CustomerRepository.
func (c *customerRepository) Save(customer model.Customer) error {
	_, err := c.db.Exec("INSERT INTO customer VALUES ($1, $2, $3, $4)", customer.Id, customer.Name, customer.PhoneNumber, customer.Address)
	if err != nil {
		fmt.Println("error from db")
		return err
	}
	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
