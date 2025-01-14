package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"user-service/db"
	"user-service/handlers"
	"user-service/services"
	"user-service/storage"
)

func main() {
	// Database connection
	database := db.InitPostgres()
	defer database.Close()

	storageService := storage.NewStorage(database)
	userService := services.NewUserService(storageService)
	handler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUsers)
		r.Post("/", handler.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetUser)
			r.Put("/", handler.UpdateUser)
			r.Delete("/", handler.DeleteUser)
		})
	})

	log.Println("User service is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
