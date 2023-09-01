package model

import (
	"time"
)

type Bill struct {
	Id          string
	BillDate    time.Time
	EntryDate   time.Time
	FinishDate  time.Time
	EmployeeId  string
	CustomerId  string
	BillDetails []BillDetail
}

type BillDetail struct {
	Id           string
	BillId       string
	Product      Product
	ProductPrice int
	Qty          int
}
