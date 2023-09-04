package delivery

import (
	"fmt"
	"go-laundry/config"
	"go-laundry/delivery/controller"
	"go-laundry/repository"
	"go-laundry/usecase"
	"os"
)

type Console struct {
	uomUC      usecase.UomUseCase
	productUC  usecase.ProductUseCase
	customerUC usecase.CustomerUseCase
	employeeUC usecase.EmployeeUseCase
}

func (c *Console) showMainMenu() {
	fmt.Println("========== ENIGMA LAUNDRY ==========\v")
	fmt.Println("1. Master Uom")
	fmt.Println("2. Master Product")
	fmt.Println("3. Master Customer")
	fmt.Println("4. Master Employee")
	fmt.Println("5. Transaksi")
	fmt.Print("6. Keluar\n\n")
	fmt.Print("Pilih Menu (1-6): ")
}

func (c *Console) Run() {
	for {
		c.showMainMenu()

		var selectedMenu string
		fmt.Scanln(&selectedMenu)

		switch selectedMenu {
		case "1":
			controller.NewUomController(c.uomUC).UomMenu()
		case "2":
			controller.NewProductController(c.productUC).ProductMenu()
		case "3":
			// controller.New
		case "4":
		case "5":
		case "6":
			os.Exit(0)
		}
	}
}

func NewConsole() *Console {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()

	// Instance Repo
	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)

	// Instance UC
	uomUC := usecase.NewUomUseCase(uomRepo)
	productUC := usecase.NewProductUseCase(productRepo, uomUC)
	customerUC := usecase.NewCustomerUseCase(customerRepo)
	employeeUC := usecase.NewEmployeeUseCase(employeeRepo)

	return &Console{
		uomUC:      uomUC,
		productUC:  productUC,
		customerUC: customerUC,
		employeeUC: employeeUC,
	}
}
