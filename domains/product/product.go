package product

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/domains/category"
	"github.com/edwardsuwirya/simpleSql/utils/appSql"
	guuid "github.com/google/uuid"
	"log"
	"strings"
)

type Product struct {
	ProductId       string            `json:"productId"`
	ProductCode     string            `json:"productCode"`
	ProductName     string            `json:"productName"`
	ProductCategory category.Category `json:"category"`
}

func AllProduct(db *sql.DB, pageNo, totalPerPage int) ([]*Product, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		SELECT p.id,p.product_code,p.product_name,p.category_id,c.category_name
		FROM m_product p join m_category c on p.category_id = c.id 
		ORDER BY p.id 
		LIMIT ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		p := new(Product)
		//c := new(Category)
		err := rows.Scan(&p.ProductId, &p.ProductCode, &p.ProductName, &p.ProductCategory.CateogryId, &p.ProductCategory.CategoryName)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func FindProductIn(db *sql.DB, ids []string) ([]*Product, error) {
	sql := `
		SELECT id,product_code,product_name
		FROM m_product 
		WHERE product_code IN(?` + strings.Repeat(",?", len(ids)-1) + `)`
	stmt, err := db.Prepare(sql)

	params := make([]interface{}, len(ids))
	for i, v := range ids {
		params[i] = v
	}
	rows, err := stmt.Query(params...)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer stmt.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.ProductId, &p.ProductCode, &p.ProductName)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func FindProductByCode(db *sql.DB, code string) (*Product, error) {
	row := db.QueryRow(`
		SELECT p.id,p.product_code,p.product_name,p.category_id
		FROM m_product p WHERE p.product_code = ?`, code)

	p := new(Product)
	err := row.Scan(&p.ProductId, &p.ProductCode, &p.ProductName, &p.ProductCategory.CateogryId)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	return p, nil
}

func CountProduct(db *sql.DB) (int, error) {
	row := db.QueryRow(`
		SELECT count(*)
		FROM m_product`)

	var totalRow int
	err := row.Scan(&totalRow)
	if err != nil {
		log.Fatalf("%v", err)
		return -1, err
	}

	return totalRow, nil
}

func CreateProduct(db *sql.DB, product *Product) (string, error) {
	res, err := appSql.WithTransaction(db, func(tx *sql.Tx) (interface{}, error) {
		id := guuid.New()
		_, err := tx.Exec("INSERT INTO m_product(id,product_code,product_name,category_id)  VALUES(?, ?, ?, ?)", id, product.ProductCode, product.ProductName, product.ProductCategory.CateogryId)
		return id.String(), err
	})

	if err != nil {
		return "", err
	}
	return res.(string), nil

}
