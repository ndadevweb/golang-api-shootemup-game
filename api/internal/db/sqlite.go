package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitSQLite() {
	var err error
	DB, err = sql.Open("sqlite", "data/app.db")
	if err != nil {
		log.Fatalf("Error SQLite: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Ping SQLite failed: %v", err)
	}
}
