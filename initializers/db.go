package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitializeDB() {
	db := DB()
	defer db.Close()

	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS teams (
            id SERIAL PRIMARY KEY,
            owner TEXT NOT NULL,
            pokemons TEXT[] NOT NULL
        );
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialized successfully")
}
