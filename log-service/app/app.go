package app

import (
	"log-service/handlers"
	"log-service/routes"
	"log-service/services"
	"log-service/storage"
	"shared/db"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	store := storage.NewStorage(db.InitPostgres())
	service := services.NewLogService(store)
	handler := handlers.NewLogHandler(service)

	return routes.InitializeRoutes(handler)
}
