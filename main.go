package main

import (
	"fmt"
	"go-laundry/config"
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

	empRepo := repository.NewEmployeeRepository(db)
	empUC := usecase.NewEmployeeUseCase(empRepo)
	employees, err := empUC.GetByName("Caca")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(employees)
}
