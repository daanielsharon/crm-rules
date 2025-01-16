package main

import (
	"log"
	"log-service/db"
	"log-service/handlers"
	"log-service/services"
	"log-service/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	store := storage.NewStorage(db.InitPostgres())
	service := services.NewLogService(store)
	handler := handlers.NewLogHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/logs", func(r chi.Router) {
		r.Get("/", handler.GetLogs)        // GET /logs
		r.Get("/{id}", handler.GetLogByID) // GET /logs/:id
	})

	log.Println("Log service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
