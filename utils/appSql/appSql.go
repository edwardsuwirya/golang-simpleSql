package appSql

import (
	"database/sql"
	"log"
)

type TrxFn func(tx *sql.Tx) (interface{}, error)

func WithTransaction(db *sql.DB, fn TrxFn) (interface{}, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	res, err := fn(tx)
	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			return nil, rollBackErr
		}

		return nil, err
	}
	return res, tx.Commit()
}
