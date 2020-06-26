package main

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/domains/product"
	"github.com/gorilla/mux"
)

const (
	PRODUCT_MAIN_ROUTE = "/product"
)

type IAppRouter interface {
	InitRoute(rt *mux.Router, db *sql.DB)
}
type appRouter struct {
	app *SimpeSql
}

func (ar *appRouter) InitMainRouter() {
	//employeeService := employee.NewEmployeeService(ar.app.DB)
	//employeeController := employee.NewEmployeeController(employeeService, ar.app.Log, ar.app.FilePath)
	//employee.NewEmployeeRoute(employeeController, ar.app.AppRouter, EMPLOYEE_MAIN_ROUTE).InitRoute()

	routerList := []IAppRouter{
		product.NewProductRoute(PRODUCT_MAIN_ROUTE),
	}

	for _, rt := range routerList {
		rt.InitRoute(ar.app.appRouter, ar.app.db)
	}
}

func NewAppRouter(app *SimpeSql) *appRouter {
	return &appRouter{app}
}
