package routes

import (
	"log-service/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRoutes(handler *handlers.LogHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/logs", func(r chi.Router) {
		r.Get("/", handler.GetLogs) // GET /logs
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetLogById) // GET /logs/:id
		})
	})

	return r
}
