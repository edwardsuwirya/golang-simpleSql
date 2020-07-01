package main

import (
	"database/sql"
	"fmt"
	"github.com/edwardsuwirya/simpleSql/config"
	"github.com/edwardsuwirya/simpleSql/domains"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type SimpeSql struct {
	db        *sql.DB
	appRouter *mux.Router
	config    *config.Conf
}

func SimpleSqlApp(c *config.Conf) *SimpeSql {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Db.DbUser, c.Db.DbPassword, c.Db.DbHost, c.Db.DbPort, c.Db.SchemaName)
	db, err := domains.NewDbInitialization(c.Db.DbEngine, dataSourceName).InitDB()
	if err != nil {
		log.Panic(err)
	}

	mainAppRouter := mux.NewRouter()

	return &SimpeSql{db: db, appRouter: mainAppRouter, config: c}
}

func (ssa *SimpeSql) run() {
	NewAppRouter(ssa).InitMainRouter()
	hostListen := fmt.Sprintf("%v:%v", ssa.config.Http.Host, ssa.config.Http.Port)
	log.Printf("Ready to listen on %v", hostListen)
	if err := http.ListenAndServe(hostListen, ssa.appRouter); err != nil {
		log.Panic(err)
	}
}

func main() {
	conf := config.NewAppConfig()
	SimpleSqlApp(conf).run()
}
