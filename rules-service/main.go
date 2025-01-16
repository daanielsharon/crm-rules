package main

import (
	"log"
	"net/http"
	"rules/db"
	"rules/handlers"
	"rules/services"
	"rules/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Database connection
	database := db.InitPostgres()
	defer database.Close()

	storageService := storage.NewStorage(database)
	ruleService := services.NewRuleService(storageService)
	handler := handlers.NewRuleHandler(ruleService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/rules", func(r chi.Router) {
		r.Get("/", handler.GetAllRulesHandler)
		r.Post("/", handler.CreateRuleHandler)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetRuleById)
			r.Put("/", handler.UpdateRuleHandler)
			r.Delete("/", handler.DeleteRuleHandler)
		})
	})

	log.Println("Rule service is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
