package domains

import (
	"database/sql"
	"fmt"
	"github.com/edwardsuwirya/simpleSql/config"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDB(c *config.Conf) (*sql.DB, error) {
	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Db.DbUser, c.Db.DbPassword, c.Db.DbHost, c.Db.DbPort, c.Db.SchemaName)
	//db, err := sql.Open("mysql", dataSourceName)

	//Alternative Way to construct Database Connection String
	cfg := &mysql.Config{
		User:   c.Db.DbUser,
		Passwd: c.Db.DbPassword,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:%v", c.Db.DbHost, c.Db.DbPort),
		DBName: c.Db.SchemaName,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panic(err)
	}

	//Ping = check database availability
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db, nil
}
