package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:admin123@tcp(localhost:3306)/test")
	return
}
