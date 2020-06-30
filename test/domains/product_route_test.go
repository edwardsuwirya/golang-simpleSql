package domains

import (
	"github.com/edwardsuwirya/simpleSql/domains/product"
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductRoute(t *testing.T) {
	t.Log("It should success init route with different prefix")
	db := myTest.DbPrep()
	defer db.Close()
	a := mux.NewRouter()
	product.NewProductRoute("/product-test").InitRoute(a, db)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/product-test", nil)
	a.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, 200)
}
