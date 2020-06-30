package domains

import (
	"github.com/edwardsuwirya/simpleSql/domains/product"
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductRoute(t *testing.T) {
	t.Log("It should return 200 when GET /product")
	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}
	db := myTest.DbPrep()
	defer db.Close()
	rr := httptest.NewRecorder()
	productControllerTest := product.NewProductController(db)
	handler := http.HandlerFunc(productControllerTest.GetProduct())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
}

func TestGetProductByCodeRoute(t *testing.T) {
	t.Log("It should return 200 when GET /product with query parameter Code")
	req, err := http.NewRequest("GET", "/product", nil)
	q := req.URL.Query()
	q.Add("code", "234")
	if err != nil {
		t.Fatal(err)
	}
	db := myTest.DbPrep()
	defer db.Close()
	rr := httptest.NewRecorder()
	productControllerTest := product.NewProductController(db)
	handler := http.HandlerFunc(productControllerTest.GetProduct())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
}
