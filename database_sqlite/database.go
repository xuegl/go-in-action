package database_sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenSqlite(dsn string) error {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}

func CloseSqlite() error {
	if DB == nil {
		return nil
	}

	err := DB.Close()
	DB = nil
	return err
}
