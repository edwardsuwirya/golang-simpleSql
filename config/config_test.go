package config

import (
	"github.com/edwardsuwirya/simpleSql/config"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestEmptyConfig(t *testing.T) {
	configTest := config.NewAppConfig()
	assert.Equal(t, configTest.Db.DbUser, "root")
	assert.Equal(t, configTest.Db.DbPassword, "password")
}

func TestConfig(t *testing.T) {
	t.Log("It should set proper configuration")
	os.Setenv("DB_USER", "TEST_USER")
	os.Setenv("DB_PASSWORD", "TEST_PASSWORD")
	os.Setenv("DB_HOST", "TEST_HOST")
	os.Setenv("DB_PORT", "TEST_PORT")
	os.Setenv("DB_SCHEMA", "TEST_SCHEMA")
	configTest := config.NewAppConfig()
	assert.Equal(t, configTest.Db.DbUser, "TEST_USER")
	assert.Equal(t, configTest.Db.DbHost, "TEST_HOST")
}
