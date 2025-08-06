package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	dsn := "root:qwertyroot@tcp(127.0.0.1:3306)/clichatdb"
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.Ping()
}
