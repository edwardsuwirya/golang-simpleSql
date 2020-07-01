package product

import (
	"fmt"
	"github.com/edwardsuwirya/simpleSql/utils/appStatus"
	myHttp "github.com/edwardsuwirya/simpleSql/utils/http"
	errorMessage "github.com/edwardsuwirya/simpleSql/utils/message"
	"net/http"
)

type ProductController struct {
	productService IProductService
	parser         *myHttp.Parser
	responder      myHttp.IResponder
}

func NewProductController(productService IProductService) *ProductController {
	return &ProductController{productService, myHttp.NewParser(), myHttp.NewDefaultJSONResponder()}
}

func (pc *ProductController) NewProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		prod := new(Product)

		pc.parser.Form(r, prod)
		newProduct, err := pc.productService.CreateAProduct(prod)
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

		fmt.Println("=======")
		fmt.Println(keys)
		var prod interface{}
		if len(keys) >= 1 {
			prod = pc.productService.GetSingleProduct(keys[0], queries[keys[0]][0])
		} else {
			//resp.Error(w, http.StatusOK, errorMessage.NewErrorMessage(appStatus.StatusNotYetImplemented))
			prod = pc.productService.GetProducts()
		}
		pc.responder.Data(w, http.StatusOK, appStatus.StatusText(0), prod)

	}
}
