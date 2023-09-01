package controller

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"
)

type UomController struct {
	uomUC usecase.UomUseCase
}

func (u *UomController) UomMenuForm() {
	fmt.Println(`
	|		+++++ Master UOM +++++	|
	| 1. Tambah Data				|
	| 2. Lihat Data					|
	| 3. Update Data				|
	| 4. Hapus Data					|
	| 5. Cari Data Berdasarkan Nama	|
	| 6. Keluar                     |
	`)
	fmt.Println("Pilih Menu (1-6): ")
	var selectMenuUom string
	fmt.Scanln(&selectMenuUom)
	switch selectMenuUom {
	case "1":
		u.insertFormUom()
	case "2":
		u.showListFormUom()
	case "3":
	case "4":
	case "5":
	case "6":
		return
	}
}

func (u *UomController) insertFormUom() {
	var uom model.Uom
	fmt.Println("Inputkan Id")
	fmt.Scanln(&uom.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&uom.Name)
	err := u.uomUC.CreateNew(uom)
	if err != nil {
		fmt.Println(err)
		return
	}
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
		fmt.Printf("ID: %s, Name: %s\n", uom.Id, uom.Name)
	}
}

func NewUomController(uomUC usecase.UomUseCase) *UomController {
	return &UomController{uomUC: uomUC}
}
