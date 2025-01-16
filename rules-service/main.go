package main

import (
	"log"
	"net/http"
	"rules-service/db"
	"rules-service/handlers"
	"rules-service/services"
	"rules-service/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Database connection
	database := db.InitPostgres()
	defer database.Close()

	ruleStorage := storage.NewRuleStorage(database)
	ruleService := services.NewRuleService(ruleStorage)
	handler := handlers.NewRuleHandler(ruleService)

	actionStorage := storage.NewActionStorage(database)
	actionService := services.NewActionService(actionStorage)
	actionHandler := handlers.NewActionHandler(actionService)

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

	r.Route("/actions", func(r chi.Router) {
		r.Post("/", actionHandler.CreateActionHandler)
	})

	log.Println("Rule service is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
