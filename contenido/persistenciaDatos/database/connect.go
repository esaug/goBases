package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connection() (*sql.DB, error) {

	// .ENV
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST "),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Open DB
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	// Verifying data conection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connection successfully")

	return db, nil
}
