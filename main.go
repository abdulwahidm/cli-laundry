package main

import (
	"fmt"
	"go-laundry/config"
	"go-laundry/model"
	"go-laundry/repository"
	"go-laundry/usecase"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()

	cstRepo := repository.NewCustomerRepository(db)
	cstUC := usecase.NewCustomerUseCase(cstRepo)
	err = cstUC.CreateNew(model.Customer{
		Id:          "3",
		Name:        "Andi",
		PhoneNumber: "0868687907",
		Address:     "Jakarta",
	})
	if err != nil {
		fmt.Println(err)
	}
}
