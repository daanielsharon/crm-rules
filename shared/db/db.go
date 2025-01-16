package db

import (
	"database/sql"
	"log"
	"shared/config"
	"shared/helpers"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitPostgres() *sql.DB {
	cfg, err := config.Load()
	helpers.PanicIfError(err)

	db, err := sql.Open("postgres", cfg.Postgres.URL)
	helpers.PanicIfError(err)

	err = db.Ping()
	helpers.PanicIfError(err)

	log.Println("Successfully connected to PostgreSQL!")
	return db
}
