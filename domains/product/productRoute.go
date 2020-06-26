package product

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type ProductRoute struct {
	prefix string
}

func (pr *ProductRoute) InitRoute(router *mux.Router, db *sql.DB) {
	//r.Use(middleware.TokenValidator)
	//r.HandleFunc("/with-category", pr.Controller.ProductWithCategory())
	//r.HandleFunc("/", pr.Controller.AllProduct()).Methods(http.MethodGet)
	productController := NewProductController(db)
	p := router.PathPrefix(pr.prefix).Subrouter()
	p.HandleFunc("", productController.GetProduct()).Methods(http.MethodGet)
	p.HandleFunc("", productController.NewProduct()).Methods(http.MethodPost)
	//p.HandleFunc("/count", productController.GetProduct()).Methods(http.MethodGet)
	//p.HandleFunc("/count-Per-Category", productController.GetProduct()).Methods(http.MethodGet)

}

func NewProductRoute(prefix string) *ProductRoute {
	return &ProductRoute{
		prefix,
	}
}
