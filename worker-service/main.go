package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"worker-service/scheduler"
	"worker-service/storage"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	db, err := sql.Open("postgres", getDatabaseURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	store := storage.NewStorage(db)
	cron := scheduler.NewScheduler(store)
	cron.Start()

	// Block indefinitely to keep the process running.
	select {}
}

func getDatabaseURL() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
