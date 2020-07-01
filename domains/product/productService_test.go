package product

import (
	"github.com/edwardsuwirya/simpleSql/domains/category"
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProductService(t *testing.T) {
	t.Log("It should create a new product")
	db := myTest.DbPrep()
	productServiceTest := NewProductService(db)
	defer db.Close()
	newProductTest := Product{
		ProductId:   "TEST_PRODUCT_ID",
		ProductCode: "TEST_PRODUCT_CODE",
		ProductName: "TEST_PRODUCT_NAME",
		ProductCategory: category.Category{
			CateogryId:   "TEST_CATEGORY_ID",
			CategoryName: "TEST_CATEGORY_NAME",
		},
	}
	p, err := productServiceTest.CreateAProduct(&newProductTest)
	if err != nil {
		t.Error("Failed Creating Product")
	}
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestGetProductService(t *testing.T) {
	t.Log("It should get list of product")
	db := myTest.DbPrep()
	productServiceTest := NewProductService(db)
	defer db.Close()
	productListTest := productServiceTest.GetProducts()
	assert.Len(t, productListTest, 2)
}
