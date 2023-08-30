package main

import (
	"fmt"
	"go-laundry/config"
)

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	Address     string
}

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

	customer := Customer{
		Id:          "1",
		Name:        "Budi",
		PhoneNumber: "081234567",
		Address:     "JL. H. Dahlan 75",
	}

	_, err = db.Exec("INSERT INTO customer VALUES ($1, $2, $3, $4)",
		customer.Id,
		customer.Name,
		customer.PhoneNumber,
		customer.Address,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Succes Inserting data")

}
