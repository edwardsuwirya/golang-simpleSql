package utils

import (
	"database/sql"
	myTest "github.com/edwardsuwirya/simpleSql/test"
	"github.com/edwardsuwirya/simpleSql/utils/appSql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockTransactional(tx *sql.Tx) (interface{}, error) {
	_, err := tx.Exec("INSERT INTO m_product(id,product_code,product_name,category_id)  VALUES('test', 'test', 'test', 'test')")
	return "test", err
}

func mockTransactionalRollback(tx *sql.Tx) (interface{}, error) {
	_, err := tx.Exec("INSERT INTO table_not_found(id,product_code,product_name,category_id)  VALUES('test', 'test', 'test', 'test')")
	return "", err
}

func TestWithTransaction(t *testing.T) {
	t.Log("It should commit transaction")
	db := myTest.DbPrep()
	defer db.Close()
	res, err := appSql.WithTransaction(db, mockTransactional)
	assert.Nil(t, err)
	assert.Equal(t, res, "test")
}

func TestWithTransactionError(t *testing.T) {
	t.Log("It should rollback transaction")
	db := myTest.DbPrep()
	defer db.Close()
	_, err := appSql.WithTransaction(db, mockTransactionalRollback)
	assert.NotNil(t, err)
}
