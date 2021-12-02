package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBOpen(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
