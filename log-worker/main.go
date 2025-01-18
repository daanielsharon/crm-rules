package main

import (
	"log"
	"log-worker/app"

	_ "github.com/lib/pq"
)

func main() {
	app.New()
	log.Println("Log worker running...")
}
