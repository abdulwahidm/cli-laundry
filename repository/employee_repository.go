package repository

import (
	"database/sql"
	"fmt"
	"go-laundry/model"
)

type EmployeeRepository interface {
	Save(employee model.Employee) error
	FindAll() ([]model.Employee, error)
	FindById(id string) (model.Employee, error)
	FindByName(name string) ([]model.Employee, error)
	Update(employee model.Employee) error
	DeleteById(id string) error
}

type employeeRepository struct {
	db *sql.DB
}

// DeleteById implements EmployeeRepository.
func (e *employeeRepository) DeleteById(id string) error {
	_, err := e.db.Exec("DELETE FROM employee WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements EmployeeRepository.
func (e *employeeRepository) FindAll() ([]model.Employee, error) {
	rows, err := e.db.Query("SELECT * FROM employee")
	if err != nil {
		return nil, err
	}
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.PhoneNumber,
			&employee.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

// FindById implements EmployeeRepository.
func (e *employeeRepository) FindById(id string) (model.Employee, error) {
	row := e.db.QueryRow("SELECT * FROM employee WHERE id = $1", id)
	var employee model.Employee
	err := row.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

// FindByName implements EmployeeRepository.
func (e *employeeRepository) FindByName(name string) ([]model.Employee, error) {
	rows, err := e.db.Query(`SELECT * FROM employee AS e WHERE e.name ILIKE $1`, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	var employees []model.Employee
	for rows.Next() {
		employee := model.Employee{}
		err := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.PhoneNumber,
			&employee.Address,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

// Save implements EmployeeRepository.
func (e *employeeRepository) Save(employee model.Employee) error {
	_, err := e.db.Exec("INSERT INTO employee VALUES ($1, $2, $3, $4)", employee.Id, employee.Name, employee.PhoneNumber, employee.Address)
	if err != nil {
		fmt.Println("error from db")
		return err
	}
	return nil
}

// Update implements EmployeeRepository.
func (e *employeeRepository) Update(employee model.Employee) error {
	_, err := e.db.Exec("UPDATE employee SET name = $2, phone_number = $3,address = $4 WHERE id = $1",
		employee.Id,
		employee.Name,
		employee.PhoneNumber,
		employee.Address,
	)
	if err != nil {
		return err
	}
	fmt.Println("update employee")
	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
