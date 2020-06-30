package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func DbPrep() *sql.DB {
	database, _ := sql.Open("sqlite3", ":memory:")
	database.Exec("CREATE TABLE IF NOT EXISTS m_product (id TEXT PRIMARY KEY, product_code TEXT, product_name TEXT,category_id TEXT)")
	database.Exec("CREATE TABLE IF NOT EXISTS m_category (id TEXT PRIMARY KEY, category_name TEXT)")

	database.Exec("INSERT INTO m_product VALUES ('111-222', '123', 'kecap','C001')")
	database.Exec("INSERT INTO m_product VALUES ('112-223', '234', 'gula','C001')")

	database.Exec("INSERT INTO m_category VALUES ('C001', 'Bumbu Dapur')")
	return database
}
