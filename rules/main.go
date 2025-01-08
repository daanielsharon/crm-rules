package main

import (
	"log"
	"net/http"
	"rules/handlers"
	"rules/routes"
	"rules/services"
	"rules/storage"
)

func main() {
	// Initialize the database
	storageLayer := storage.NewStorage("rules.db")

	// Initialize services
	service := services.NewRuleService(storageLayer)
	handler := handlers.NewRuleHandler(service)

	// Initialize routes
	router := routes.InitializeRoutes(handler)

	// Start the server
	log.Println("Rules Service is running on port 8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
