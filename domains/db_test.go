package domains

import (
	"github.com/edwardsuwirya/simpleSql/config"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	os.Setenv("DB_ENGINE", "sqlite3")
	configTest := config.NewAppConfig()
	dbInitTest := NewDbInitialization(configTest.Db.DbEngine, ":memory:")
	dbConn, err := dbInitTest.InitDB()

	assert.Nil(t, err)
	assert.NotNil(t, dbConn)

}
