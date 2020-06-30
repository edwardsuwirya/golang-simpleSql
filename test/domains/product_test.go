package domains

import (
	"github.com/edwardsuwirya/simpleSql/domains/product"
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllProduct(t *testing.T) {
	t.Log("It should return all product")
	db := myTest.DbPrep()
	defer db.Close()
	result, _ := product.AllProduct(db, 0, 1000)
	assert.Equal(t, len(result), 2)
}

func TestFindProductIn(t *testing.T) {
	t.Log("It should return all product from collection of product code")
	db := myTest.DbPrep()
	defer db.Close()
	result, _ := product.FindProductIn(db, []string{"123"})
	assert.Equal(t, len(result), 1)
}

func TestFindProductInNoRow(t *testing.T) {
	t.Log("It should return all product from collection of product code")
	db := myTest.DbPrep()
	defer db.Close()
	result, _ := product.FindProductIn(db, []string{"XXX"})
	assert.Equal(t, len(result), 0)
}
