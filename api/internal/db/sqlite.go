package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import side-effect pour le driver SQLite
)

var db *sql.DB // Instance partagée de la base SQLite

// InitSQLite initialise la connexion à SQLite et la stocke globalement
func InitSQLite() {
	var err error

	// Connexion à la base de données SQLite (fichier situé dans data/app.db)
	db, err = sql.Open("sqlite", "data/app.db")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de SQLite : %v", err)
	}

	// Vérification de la connexion
	if err := db.Ping(); err != nil {
		log.Fatalf("Erreur de connexion à SQLite : %v", err)
	}
}

// GetDB retourne l'instance de base de données SQLite
func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Base SQLite non initialisée : appelle InitSQLite() avant.")
	}
	return db
}
