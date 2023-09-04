package controller

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"
)

type UomController struct {
	uomUC usecase.UomUseCase
}

func (u *UomController) UomMenu() {
	fmt.Println("========== Master UOM ==========\v")
	fmt.Println("1. Add New Uom")
	fmt.Println("2. View All Uom")
	fmt.Println("3. Update Uom ")
	fmt.Println("4. Delete Uom")
	fmt.Println("5. Find By Name")
	fmt.Print("6. Comeback to Main Menu\n\n")
	fmt.Print("Choose Menu (1-6): ")

	var selectMenuUom string
	fmt.Scanln(&selectMenuUom)

	switch selectMenuUom {
	case "1":
		u.insertFormUom()
	case "2":
		u.showListFormUom()
	case "3":
		u.updateFormUom()
	case "4":
		u.DeleteFormUom()
	case "5":
		u.FindByNameFormUom()
	case "6":
		return
	}
}

func (u *UomController) insertFormUom() {
	var uom model.Uom
	fmt.Print("Input Id: ")
	fmt.Scanln(&uom.Id)
	fmt.Print("Input Name: ")
	fmt.Scanln(&uom.Name)
	err := u.uomUC.CreateNew(uom)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("New uom has been added..")
}

func (u *UomController) showListFormUom() {
	uoms, err := u.uomUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(uoms) == 0 {
		fmt.Println("Uom is empty")
		return
	}
	for _, uom := range uoms {
		fmt.Println("ID :", uom.Id, "Name : ", uom.Name)
	}
}

func (u *UomController) updateFormUom() {
	var uom model.Uom

	fmt.Print("Input Id: ")
	fmt.Scanln(&uom.Id)

	fmt.Print("Input Name: ")
	fmt.Scanln(&uom.Name)

	err := u.uomUC.Update(uom)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Uom with id %s, succsessfully updated!\n", uom.Id)
}

func (u *UomController) FindByNameFormUom() {
	var name string
	fmt.Print("Input searching name uom: ")
	fmt.Scanln(&name)

	uoms, err := u.uomUC.FindByName(name)
	if err != nil {
		fmt.Println(err)
	}
	for _, uom := range uoms {
		fmt.Println("ID :", uom.Id, "Name : ", uom.Name)
	}

}

func (u *UomController) DeleteFormUom() {
	var id string
	fmt.Printf("Input id to deleted: ")
	fmt.Scanln(&id)
	err := u.uomUC.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Uom with id %s has been deleted\n", id)
}

func NewUomController(uomUC usecase.UomUseCase) *UomController {
	return &UomController{uomUC: uomUC}
}
