package product

import (
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllProduct(t *testing.T) {
	t.Run("It should return all product", func(t *testing.T) {
		db := myTest.DbPrep()
		defer db.Close()
		result, _ := AllProduct(db, 0, 1000)
		assert.Equal(t, len(result), 2)
	})
}

func TestFindProductIn(t *testing.T) {
	t.Run("It should return all product from collection of product code", func(t *testing.T) {
		db := myTest.DbPrep()
		defer db.Close()
		result, _ := FindProductIn(db, []string{"123"})
		assert.Equal(t, len(result), 1)
	})
}

func TestFindProductInNoRow(t *testing.T) {
	t.Run("It should return all product from collection of product code", func(t *testing.T) {
		db := myTest.DbPrep()
		defer db.Close()
		result, _ := FindProductIn(db, []string{"XXX"})
		assert.Equal(t, len(result), 0)
	})
}
