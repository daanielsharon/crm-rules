package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Println("Running migrations...")
	runMigrations()

	log.Println("Seeding initial data...")
	seedData()

	log.Println("Migration and seeding completed successfully!")
}

// runMigrations executes the migrations using `golang-migrate`
func runMigrations() {
	dbURL := getDatabaseURL()
	migrationsPath := "./migrations"

	cmd := exec.Command("migrate", "-database", dbURL, "-path", migrationsPath, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func seedData() {
	dbURL := getDatabaseURL()
	seedFile := "./migrations/seed_rules_table.up.sql"

	cmd := exec.Command("psql", dbURL, "-f", seedFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}
}

func getDatabaseURL() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}
