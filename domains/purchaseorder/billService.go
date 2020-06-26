package purchaseorder

import (
	"database/sql"
)

type BillService struct {
	db *sql.DB
}

func NewBillService(db *sql.DB) *BillService {
	return &BillService{db}
}

func (bs *BillService) CreateABill(billId int, productId int, sales float64, tax float64) *Bill {
	bill := Bill{
		BillId:    billId,
		ProductId: productId,
		Sales:     sales,
		Tax:       tax,
	}
	err := CreateBill(bs.db, bill)
	if err != nil {
		return nil
	}
	return &bill
}
func (bs *BillService) GetAllBill() []*Bill {
	bills, err := AllBill(bs.db)
	if err != nil {
		return nil
	}
	return bills
}

func (bs *BillService) TotalSales() float64 {
	bills, err := TotalSales(bs.db)
	if err != nil {
		return 0
	}
	return (*bills).Summary

}
