package main

import (
	"database/sql"
	"fmt"
	"os"

	"shared/config"
	"shared/db"
	"shared/helpers"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations(cfg *config.Config) {
	m, err := migrate.New(
		"file://./migrations",
		cfg.Postgres.URL,
	)
	helpers.PanicIfError(err)

	err = m.Up()
	helpers.PanicIfError(err)

	fmt.Println("Migration applied successfully!")
}

func seedData() {
	seedFiles, err := os.ReadDir("./seeds")
	helpers.PanicIfError(err)

	db := db.InitPostgres()
	for _, seedFile := range seedFiles {
		filePath := fmt.Sprintf("./seeds/%s", seedFile.Name())
		err := executeSQLFile(db, filePath)
		helpers.PanicIfError(err)
	}

	fmt.Println("Seeding applied successfully!")
}

func executeSQLFile(db *sql.DB, filePath string) error {
	sqlQuery, err := os.ReadFile(filePath)
	helpers.PanicIfError(err)

	_, err = db.Exec(string(sqlQuery))
	helpers.PanicIfError(err)

	return nil
}
