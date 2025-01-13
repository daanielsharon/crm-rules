package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func main() {
	log.Println("Running migrations...")
	runMigrations()

	log.Println("Seeding initial data...")
	seedData()

	log.Println("Migration and seeding completed successfully!")
}

func runMigrations() {
	dbURL := getDatabaseURL()
	fmt.Println("Database URL:", dbURL)

	m, err := migrate.New(
		"file://./migrations",
		dbURL,
	)

	if err != nil {
		log.Fatalf("Migration error: %v", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err.Error())
	}

	fmt.Println("Migration applied successfully!")
}

func seedData() {
	dbURL := getDatabaseURL()
	fmt.Println("Database URL:", dbURL)
	seedFiles, err := os.ReadDir("./seeds")
	if err != nil {
		log.Fatalf("Failed to read seeds directory: %v", err)
	}

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	for _, seedFile := range seedFiles {
		filePath := fmt.Sprintf("./seeds/%s", seedFile.Name())
		err := executeSQLFile(db, filePath)
		if err != nil {
			log.Fatalf("failed to execute seed file %s: %v", seedFile.Name(), err)
		}
	}

	fmt.Println("Seeding applied successfully!")
}

func executeSQLFile(db *sqlx.DB, filePath string) error {
	sqlQuery, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file %s: %w", filePath, err)
	}

	_, err = db.Exec(string(sqlQuery))
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}
	return nil
}

func getDatabaseURL() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}
