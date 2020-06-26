package purchaseorder

import (
	"database/sql"
	"log"
)

type TotalBill struct {
	Summary float64
}

type Bill struct {
	BillId    int
	ProductId int
	Sales     float64
	Tax       float64
}

func TotalSales(db *sql.DB) (*TotalBill, error) {
	rows, err := db.Query("SELECT sum(sales) as total FROM bill")
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	totalSales := new(TotalBill)

	for rows.Next() {
		if err := rows.Scan(&totalSales.Summary); err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
	}
	return totalSales, nil
}
func AllBill(db *sql.DB) ([]*Bill, error) {
	rows, err := db.Query("SELECT * FROM bill")
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	bills := make([]*Bill, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type
		b := new(Bill)
		err := rows.Scan(&b.BillId, &b.ProductId, &b.Sales, &b.Tax)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		bills = append(bills, b)
	}

	return bills, nil

}

func CreateBill(db *sql.DB, bill Bill) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	stmt, err := db.Prepare("INSERT INTO bill  VALUES(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(bill.BillId, bill.ProductId, bill.Sales, bill.Tax); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}
