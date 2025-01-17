package app

import (
	"shared/db"
	"user-service/handlers"
	"user-service/route"
	"user-service/services"
	"user-service/storage"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	database := db.InitPostgres()

	storageService := storage.NewStorage(database)
	userService := services.NewUserService(storageService)
	handler := handlers.NewUserHandler(userService)

	return route.InitializeRoutes(handler)
}
