package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testee")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}
}
