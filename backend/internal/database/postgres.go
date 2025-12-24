package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Init initialise la connexion à PostgreSQL
func Init() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "host=localhost port=5432 user=postgres password=password dbname=quizdb sslmode=disable"
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Impossible de ping la base de données:", err)
	}

	log.Println("Connexion à PostgreSQL réussie")
}

// Close ferme la connexion à la base de données
func Close() {
	if DB != nil {
		DB.Close()
	}
}
