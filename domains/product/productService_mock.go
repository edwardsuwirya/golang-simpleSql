package product

import (
	"github.com/edwardsuwirya/simpleSql/domains/productPrice"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m ProductServiceMock) GetProductsPaging(pageNo, totalPerPage int) []*Product {
	panic("implement me")
}

func (m ProductServiceMock) GetProducts() []*Product {
	ret := m.Called()
	var productResultMock = make([]*Product, 0)
	if ret.Get(0) != nil {
		productResultMock = ret.Get(0).([]*Product)
	}
	return productResultMock
}

func (m ProductServiceMock) GetSingleProduct(field string, code string) *Product {
	ret := m.Called(field, code)
	var productResultMock *Product
	if ret.Get(0) != nil {
		productResultMock = ret.Get(0).(*Product)
	}
	return productResultMock
}

func (m ProductServiceMock) GetProductsIn(ids []string) []*Product {
	panic("implement me")
}

func (m ProductServiceMock) CreateAProduct(newProduct *Product) (*Product, error) {
	panic("implement me")
}

func (m ProductServiceMock) GetProductWithPrice() []*productPrice.ProductPrice {
	panic("implement me")
}
