package config

import (
	"github.com/edwardsuwirya/simpleSql/utils"
)

type httpConf struct {
	Host string
	Port string
}
type dbConf struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	SchemaName string
	DbEngine   string
}

type Conf struct {
	Db   dbConf
	Http httpConf
}

func NewAppConfig() *Conf {
	return &Conf{
		dbConf{
			DbUser:     utils.ViperGetEnv("DB_USER", "root"),
			DbPassword: utils.ViperGetEnv("DB_PASSWORD", "password"),
			DbHost:     utils.ViperGetEnv("DB_HOST", "localhost"),
			DbPort:     utils.ViperGetEnv("DB_PORT", "3306"),
			SchemaName: utils.ViperGetEnv("DB_SCHEMA", "schema"),
			DbEngine:   utils.ViperGetEnv("DB_ENGINE", "mysql"),
		},
		httpConf{Host: utils.ViperGetEnv("APP_HOST", ""), Port: utils.ViperGetEnv("APP_PORT", "8080")}}
}
