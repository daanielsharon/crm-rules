package main

import (
	"log"
	"net/http"

	"user-service/app"
)

func main() {
	log.Println("User service is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", app.New()))
}
