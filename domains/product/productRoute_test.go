package product

import (
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductRoute(t *testing.T) {
	t.Run("It should success init route with different prefix", func(t *testing.T) {
		db := myTest.DbPrep()
		defer db.Close()
		a := mux.NewRouter()
		NewProductRoute("/product-test").InitRoute(a, db)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/product-test", nil)
		a.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 200)
	})
}
