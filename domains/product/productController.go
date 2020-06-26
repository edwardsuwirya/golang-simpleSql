package product

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/utils/appStatus"
	myHttp "github.com/edwardsuwirya/simpleSql/utils/http"
	errorMessage "github.com/edwardsuwirya/simpleSql/utils/message"
	"net/http"
)

type ProductController struct {
	prodService *ProductService
	parser      *myHttp.Parser
	responder   myHttp.IResponder
}

func NewProductController(db *sql.DB) *ProductController {
	prodService := NewProductService(db)
	return &ProductController{prodService, myHttp.NewParser(), myHttp.NewDefaultJSONResponder()}
}

func (pc *ProductController) NewProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		prod := new(Product)

		pc.parser.Form(r, prod)
		newProduct, err := pc.prodService.CreateAProduct(prod)
		if err != nil {
			pc.responder.Error(w, http.StatusInternalServerError, errorMessage.NewErrorMessage(appStatus.Error))
		}
		pc.responder.Data(w, http.StatusOK, appStatus.StatusText(0), newProduct)
	}
}

func (pc *ProductController) GetProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		keys := []string{}
		for k := range queries {
			keys = append(keys, k)
		}

		var prod interface{}
		if len(keys) > 0 {
			prod = pc.prodService.GetSingleProduct(keys[0], queries[keys[0]][0])
		} else {
			//resp.Error(w, http.StatusOK, errorMessage.NewErrorMessage(appStatus.StatusNotYetImplemented))
			prod = pc.prodService.GetProducts()
		}
		pc.responder.Data(w, http.StatusOK, appStatus.StatusText(0), prod)

	}
}
