package controller

import (
	"bufio"
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"
	"os"
	"strconv"
)

type ProductController struct {
	productUC usecase.ProductUseCase
}

func (u *ProductController) ProductMenu() {
	fmt.Println("========== Master Product ==========\v")
	fmt.Println("1. Add New Product")
	fmt.Println("2. View All Product")
	fmt.Println("3. Update Product ")
	fmt.Println("4. Delete Product")
	fmt.Println("5. Find By Name")
	fmt.Print("6. Comeback to Main Menu\n\n")
	fmt.Print("Choose Menu (1-6): ")

	var selectMenuProduct string
	fmt.Scanln(&selectMenuProduct)

	switch selectMenuProduct {
	case "1":
		u.insertFormProduct()
	case "2":
		u.showListFormProduct()
	case "3":
		u.updateFormProduct()
	case "4":
		u.DeleteFormProduct()
	case "5":
		u.FindByNameFormProduct()
	case "6":
		return
	}
}

func (u *ProductController) insertFormProduct() {
	var product model.Product

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Id Product: ")
	scanner.Scan()
	product.Id = scanner.Text()

	fmt.Print("Input Name Product: ")
	scanner.Scan()
	product.Name = scanner.Text()

	fmt.Print("Input Price Product: ")
	scanner.Scan()
	priceStr := scanner.Text()
	product.Price, _ = strconv.Atoi(priceStr)

	fmt.Print("Input Uom Id: ")
	scanner.Scan()
	product.Uom.Id = scanner.Text()

	err := u.productUC.CreateNew(product)
	if err != nil {
		fmt.Println(err)
	} else {
	}
	fmt.Println("New product has been added..")
}

func (u *ProductController) showListFormProduct() {
	products, err := u.productUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(products) == 0 {
		fmt.Println("product is empty")
		return
	}
	for _, product := range products {
		fmt.Println(product)
	}
}

func (u *ProductController) updateFormProduct() {
	var product model.Product
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Id Product to Update: ")
	fmt.Scanln(&product.Id)

	fmt.Print("Update Name Product: ")
	scanner.Scan()
	product.Name = scanner.Text()

	fmt.Print("Update Price Product: ")
	scanner.Scan()
	priceStr := scanner.Text()
	product.Price, _ = strconv.Atoi(priceStr)

	fmt.Print("Update Uom Id: ")
	scanner.Scan()
	product.Uom.Id = scanner.Text()

	err := u.productUC.Update(product)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("product with id %s, succsessfully updated!\n", product.Id)
}

func (u *ProductController) FindByNameFormProduct() {
	var name string
	fmt.Print("Input searching name product: ")
	fmt.Scanln(&name)

	products, err := u.productUC.GetByName(name)
	if err != nil {
		fmt.Println(err)
	}
	for _, product := range products {
		fmt.Println("ID :", product.Id, "Name : ", product.Name)
	}

}

func (u *ProductController) DeleteFormProduct() {
	var id string
	fmt.Printf("Input id to deleted: ")
	fmt.Scanln(&id)
	err := u.productUC.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("product with id %s has been deleted\n", id)
}

func NewProductController(productUC usecase.ProductUseCase) *ProductController {
	return &ProductController{productUC: productUC}
}
