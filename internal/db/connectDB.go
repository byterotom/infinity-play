package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "internal/db/infinity.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
