package product

import (
	"github.com/edwardsuwirya/simpleSql/domains/category"
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductController(t *testing.T) {
	t.Run("It should return 200 when GET /product", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/product", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		productResultMock := []*Product{
			{
				ProductId:   "111",
				ProductCode: "222",
				ProductName: "333",
				ProductCategory: category.Category{
					CateogryId:   "444",
					CategoryName: "555",
				},
			},
		}
		productServiceMock := new(ProductServiceMock)
		productServiceMock.On("GetProducts").Return(productResultMock)
		productControllerTest := NewProductController(productServiceMock)

		handler := http.HandlerFunc(productControllerTest.GetProduct())
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, 200)

	})
}

func TestGetProductByCodeController(t *testing.T) {
	t.Run("It should return 200 when GET /product with query parameter Code", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/product", nil)
		q := req.URL.Query()
		q.Add("code", "234")
		req.URL.RawQuery = q.Encode()
		if err != nil {
			t.Fatal(err)
		}
		db := myTest.DbPrep()
		defer db.Close()
		rr := httptest.NewRecorder()

		productResultMock := Product{
			ProductId:   "111",
			ProductCode: "222",
			ProductName: "333",
			ProductCategory: category.Category{
				CateogryId:   "444",
				CategoryName: "555",
			},
		}
		productServiceMock := new(ProductServiceMock)
		productServiceMock.On("GetSingleProduct", "code", "234").Return(&productResultMock)
		productControllerTest := NewProductController(productServiceMock)

		handler := http.HandlerFunc(productControllerTest.GetProduct())
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, 200)
	})
}
