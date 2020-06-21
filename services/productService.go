package services

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/models"
)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) GetProducts(pageNo, totalPerPage int) []*models.Product {
	products, err := models.AllProduct(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return products
}

func (ps *ProductService) GetProductsIn(ids []string) []*models.Product {
	products, err := models.FindProductIn(ps.db, ids)
	if err != nil {
		return nil
	}
	return products
}

func (ps *ProductService) CreateAProduct(productCode string, productName string, category string) (*models.Product, error) {
	product := &models.Product{
		ProductCode:     productCode,
		ProductName:     productName,
		ProductCategory: models.Category{CateogryId: category},
	}
	id, err := models.CreateProduct(ps.db, product)
	if err != nil {
		return nil, err
	}
	product.ProductId = id
	return product, nil
}

func (ps *ProductService) GetProductWithPrice() []*models.ProductPrice {
	productPrice, err := models.AllProductPrice(ps.db)
	if err != nil {
		return nil
	}
	return productPrice
}
