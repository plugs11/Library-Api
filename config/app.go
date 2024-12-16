package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Connect() {
	db, err = sql.Open("mysql", "root:manan@123@tcp(127.0.0.1:3306)/LIBRARY")
	if err != nil {
		fmt.Println("Error in connecting to database")
	}
}
func GetDB() *sql.DB {
	return db
}

////connecting to database
