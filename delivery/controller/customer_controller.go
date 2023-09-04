package controller

import (
	"bufio"
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"
	"os"
)

type CustomerController struct {
	customerUC usecase.CustomerUseCase
}

func (u *CustomerController) CustomerMenu() {
	fmt.Println("========== Master UOM ==========\v")
	fmt.Println("1. Add New Customer")
	fmt.Println("2. View All Customer")
	fmt.Println("3. Update Customer ")
	fmt.Println("4. Delete Customer")
	fmt.Println("5. Find By Name")
	fmt.Print("6. Comeback to Main Menu\n\n")
	fmt.Print("Choose Menu (1-6): ")

	var selectMenuCustomer string
	fmt.Scanln(&selectMenuCustomer)

	switch selectMenuCustomer {
	case "1":
		u.insertFormCustomer()
	case "2":
		u.showListFormCustomer()
	case "3":
		u.updateFormCustomer()
	case "4":
		u.DeleteFormCustomer()
	case "5":
		u.FindByNameFormCustomer()
	case "6":
		return
	}
}

func (c *CustomerController) insertFormCustomer() {
	var customer model.Customer

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Id Customer: ")
	scanner.Scan()
	customer.Id = scanner.Text()

	fmt.Print("Input Name Customer: ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Input Price Customer: ")
	scanner.Scan()
	customer.PhoneNumber = scanner.Text()

	fmt.Print("Input Addres Customer: ")
	scanner.Scan()
	customer.Address = scanner.Text()

	err := c.customerUC.CreateNew(customer)
	if err != nil {
		fmt.Println(err)
	} else {
	}
	fmt.Println("New customer has been added..")
}

func (u *CustomerController) showListFormCustomer() {
	customers, err := u.customerUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(customers) == 0 {
		fmt.Println("product is empty")
		return
	}
	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func (u *CustomerController) updateFormCustomer() {
	var customer model.Customer
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Id Customer to Update: ")
	fmt.Scanln(&customer.Id)

	fmt.Print("Update Name Customer: ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Update Phone Number Customer: ")
	scanner.Scan()
	customer.PhoneNumber = scanner.Text()

	fmt.Print("Update Address Customer: ")
	scanner.Scan()
	customer.Address = scanner.Text()

	err := u.customerUC.Update(customer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("product with id %s, succsessfully updated!\n", customer.Id)
}

func (c *CustomerController) FindByNameFormCustomer() {
	var name string
	fmt.Print("Input searching name Customer: ")
	fmt.Scanln(&name)

	customers, err := c.customerUC.GetByName(name)
	if err != nil {
		fmt.Println(err)
	}
	for _, product := range customers {
		fmt.Println("ID :", product.Id, "Name : ", product.Name)
	}

}

func (c *CustomerController) DeleteFormCustomer() {
	var id string
	fmt.Printf("Input id to deleted: ")
	fmt.Scanln(&id)
	err := c.customerUC.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("product with id %s has been deleted\n", id)
}

func NewCustomerController(customerUC usecase.CustomerUseCase) *CustomerController {
	return &CustomerController{customerUC: customerUC}
}
