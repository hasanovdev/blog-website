package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Hasanov@0301@/users")
	if err != nil {
		panic(err)
	}
	return db
}
