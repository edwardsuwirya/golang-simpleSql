package domains

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type dbInitialization struct {
	dbEngine       string
	dataSourceName string
}

func NewDbInitialization(engine, dataSource string) *dbInitialization {
	return &dbInitialization{engine, dataSource}
}

func (dbi *dbInitialization) InitDB() (*sql.DB, error) {
	db, err := sql.Open(dbi.dbEngine, dbi.dataSourceName)

	//Alternative Way to construct Database Connection String
	//cfg := &mysql.Config{
	//	User:   c.Db.DbUser,
	//	Passwd: c.Db.DbPassword,
	//	Net:    "tcp",
	//	Addr:   fmt.Sprintf("%v:%v", c.Db.DbHost, c.Db.DbPort),
	//	DBName: c.Db.SchemaName,
	//}
	//db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panic(err)
	}

	//Ping = check database availability
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db, nil
}
