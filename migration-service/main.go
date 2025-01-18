package main

import (
	"fmt"
	"shared/config"
	"shared/helpers"
)

func main() {
	cfg, err := config.Load()
	helpers.PanicIfError(err)

	runMigrations(cfg)
	seedData()

	fmt.Println("Database migration and seeding completed successfully!")
}
