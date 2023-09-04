package controller

import (
	"bufio"
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"
	"os"
)

type EmployeeController struct {
	employeeUC usecase.EmployeeUseCase
}

func (e *EmployeeController) EmployeeMenu() {
	fmt.Println("========== Master Employee ==========\v")
	fmt.Println("1. Add New Employee")
	fmt.Println("2. View All Employee")
	fmt.Println("3. Update Employee ")
	fmt.Println("4. Delete Employee")
	fmt.Println("5. Find By Name")
	fmt.Print("6. Comeback to Main Menu\n\n")
	fmt.Print("Choose Menu (1-6): ")

	var selectMenuEmployee string
	fmt.Scanln(&selectMenuEmployee)

	switch selectMenuEmployee {
	case "1":
		e.insertFormEmployee()
	case "2":
		e.showListFormEmployee()
	case "3":
		e.updateFormEmployee()
	case "4":
		e.DeleteFormEmployee()
	case "5":
		e.FindByNameFormEmployee()
	case "6":
		return
	}
}

func (e *EmployeeController) insertFormEmployee() {
	var employee model.Employee

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Id Employee: ")
	scanner.Scan()
	employee.Id = scanner.Text()

	fmt.Print("Input Name Employee: ")
	scanner.Scan()
	employee.Name = scanner.Text()

	fmt.Print("Input Price Employee: ")
	scanner.Scan()
	employee.PhoneNumber = scanner.Text()

	fmt.Print("Input Addres Employee: ")
	scanner.Scan()
	employee.Address = scanner.Text()

	err := e.employeeUC.CreateNew(employee)
	if err != nil {
		fmt.Println(err)
	} else {
	}
	fmt.Println("New employee has been added..")
}

func (e *EmployeeController) showListFormEmployee() {
	employees, err := e.employeeUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(employees) == 0 {
		fmt.Println("product is empty")
		return
	}
	for _, employee := range employees {
		fmt.Println(employee)
	}
}

func (e *EmployeeController) updateFormEmployee() {
	var employee model.Employee
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Id Customer to Update: ")
	fmt.Scanln(&employee.Id)

	fmt.Print("Update Name Employee: ")
	scanner.Scan()
	employee.Name = scanner.Text()

	fmt.Print("Update Phone Number Employee: ")
	scanner.Scan()
	employee.PhoneNumber = scanner.Text()

	fmt.Print("Update Address Customer: ")
	scanner.Scan()
	employee.Address = scanner.Text()

	err := e.employeeUC.Update(employee)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("product with id %s, succsessfully updated!\n", employee.Id)
}

func (e *EmployeeController) FindByNameFormEmployee() {
	var name string
	fmt.Print("Input searching name Employee: ")
	fmt.Scanln(&name)

	employees, err := e.employeeUC.GetByName(name)
	if err != nil {
		fmt.Println(err)
	}
	for _, employee := range employees {
		fmt.Println(employee)
	}

}

func (e *EmployeeController) DeleteFormEmployee() {
	var id string
	fmt.Printf("Input id to deleted: ")
	fmt.Scanln(&id)
	err := e.employeeUC.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("product with id %s has been deleted\n", id)
}

func NewEmployeeController(employeeUC usecase.EmployeeUseCase) *EmployeeController {
	return &EmployeeController{employeeUC: employeeUC}
}
