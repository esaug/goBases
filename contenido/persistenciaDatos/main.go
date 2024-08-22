package main

import (
	"log"
	"persistencia/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Connect to DB
	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
