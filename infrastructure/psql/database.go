package psql

import (
	"database/sql"
	"log"
	"os"
)

func GetDBCon() (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db, nil
}
