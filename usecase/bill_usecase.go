package usecase

import "go-laundry/model"

/*
1. Create Transaksi
2. History Transaksi Berdasarkan id | model.Bill
3. Get All History | []model.Bill
*/

type BillUseCase interface {
	CreateNewTransaction(payload model.Bill) error
}

type billUseCase struct {
	productUC ProductUseCase
	// employeeUC EmployeeUseCase
	// customerUC CustomerUseCase
}

// CreateNewTransaction implements BillUseCase.
func (*billUseCase) CreateNewTransaction(payload model.Bill) error {
	// Validasi Employee dari id
	// Validasi Customer dari id

	// for _, billDetail := range payload.BillDetails {
	// 	Validasi Product dari id
	// 	billDetail.Product = product
	// 	billDetail.ProductPrice = product.Price <- Reassign harga pelayanan
	// }

	// Simpan Bill panggil repo
	return nil
}

func NewBillUseCase(productUC ProductUseCase) BillUseCase {
	return &billUseCase{
		productUC: productUC,
	}
}
