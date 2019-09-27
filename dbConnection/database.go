package dbConnection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	db, _ = sql.Open("mysql", "root:abc513@tcp(localhost:3310)/stepbox")
}

func PrintVersionOfDB() {
	var name string
	err := db.QueryRow("SELECT VERSION();").Scan(&name)

	if err != nil {
		fmt.Println("Query Error")
		panic(err)
	}
	fmt.Println(name)
}

func GetDB() *sql.DB {
	return db
}