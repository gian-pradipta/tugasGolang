package database

import (
	"database/sql"

	_ "github.com/glebarez/sqlite"
)

const DB_NAME string = "test.db"
const DB_DRIVER string = "sqlite"

func New() *sql.DB {
	db, err := sql.Open(DB_DRIVER, DB_NAME)
	if err != nil {
		panic(err)
	}
	db.Exec("PRAGMA foreign_keys = ON;")
	return db
}
