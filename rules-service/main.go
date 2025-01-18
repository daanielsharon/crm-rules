package main

import (
	"log"
	"net/http"

	"rules-service/app"
)

func main() {
	log.Println("Rule service is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", app.New()))
}
