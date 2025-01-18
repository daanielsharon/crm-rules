package main

import (
	"log"
	"shared/helpers"

	"worker-service/app"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	application, err := app.New()
	helpers.PanicIfError(err)

	log.Println("Starting worker service...")
	application.Start()

	// Block indefinitely to keep the process running
	select {}
}
