package main

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/config"
	"github.com/edwardsuwirya/simpleSql/models"
	"github.com/edwardsuwirya/simpleSql/services"
	"log"
)

type simpeSql struct {
	db *sql.DB
}

func SimpleSqlApp(c *config.Conf) *simpeSql {
	db, err := models.InitDB(c)
	if err != nil {
		log.Panic(err)
	}
	return &simpeSql{db: db}
}

func (ssa *simpeSql) run() {
	//billService := services.NewBillService(env.db)
	//newBill := billService.CreateABill(8, 4, 12000, 8)
	//if newBill != nil {
	//	log.Print(*newBill)
	//}
	//sales := billService.TotalSales()
	//log.Printf("%v", humanize.Commaf(sales))

	productService := services.NewProductService(ssa.db)

	prod, err := productService.CreateAProduct("KGB", "Sofa", "5403a1a0-5520-11ea-bb2b-9378803a9e60")
	if err != nil {
		log.Print(err)
	} else {
		log.Print(*prod)
	}
	//products := productService.GetProducts(1, 2)
	//products := productService.GetProductsIn([]string{"DEA", "ZZZ"})
	//for _, p := range products {
	//	log.Printf("%v %v %v %v %v", p.ProductId, p.ProductCode, p.ProductName, p.ProductCategory.CateogryId, p.ProductCategory.CategoryName)

	//prods := productService.GetProductWithPrice()
	//for _, pp := range prods {
	//	log.Printf("%v", pp)
	//}
}

func main() {
	conf := config.NewAppConfig()
	SimpleSqlApp(conf).run()
}
