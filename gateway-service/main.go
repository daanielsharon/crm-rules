package main

import (
	"gateway/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.InitializeRoutes()

	log.Println("Gateway is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
