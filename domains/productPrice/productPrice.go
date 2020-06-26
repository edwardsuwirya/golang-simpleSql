package productPrice

import (
	"database/sql"
	"log"
)

type ProductPrice struct {
	productPriceId string
	productId      string
	productCode    string
	productName    string
	categoryName   string
	productPrice   float64
	isActive       string
}

func AllProductPrice(db *sql.DB) ([]*ProductPrice, error) {
	rows, err := db.Query("SELECT * FROM v_product_price")
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	productPrices := make([]*ProductPrice, 0)

	for rows.Next() {
		pp := new(ProductPrice)
		err := rows.Scan(&pp.productPriceId, &pp.productId, &pp.productCode, &pp.productName, &pp.categoryName, &pp.productPrice, &pp.isActive)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		productPrices = append(productPrices, pp)
	}

	return productPrices, nil

}
