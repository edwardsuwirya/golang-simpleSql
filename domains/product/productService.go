package product

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/domains/productPrice"
)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) GetProductsPaging(pageNo, totalPerPage int) []*Product {
	products, err := AllProduct(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return products
}
func (ps *ProductService) GetProducts() []*Product {
	tot, _ := CountProduct(ps.db)
	products, err := AllProduct(ps.db, 0, tot)
	if err != nil {
		return nil
	}
	return products
}
func (ps *ProductService) GetSingleProduct(field string, code string) *Product {
	var product *Product
	var err error
	switch field {
	case "code":
		product, err = FindProductByCode(ps.db, code)
	}

	if err != nil {
		return nil
	}
	return product
}
func (ps *ProductService) GetProductsIn(ids []string) []*Product {
	products, err := FindProductIn(ps.db, ids)
	if err != nil {
		return nil
	}
	return products
}

func (ps *ProductService) CreateAProduct(newProduct *Product) (*Product, error) {
	id, err := CreateProduct(ps.db, newProduct)
	if err != nil {
		return nil, err
	}
	newProduct.ProductId = id
	return newProduct, nil
}
func (ps *ProductService) GetProductWithPrice() []*productPrice.ProductPrice {
	productPrice, err := productPrice.AllProductPrice(ps.db)
	if err != nil {
		return nil
	}
	return productPrice
}
