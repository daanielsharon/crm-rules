package routes

import (
	"crm-rules/gateway/handlers"
	localMiddleware "crm-rules/gateway/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(localMiddleware.RateLimiter())

	router.Route("/rules", func(r chi.Router) {
		r.Post("/", handlers.CreateRuleHandler)    // POST /rules
		r.Put("/{id}", handlers.UpdateRuleHandler) // PUT /rules/:id
	})

	router.Get("/logs", handlers.FetchLogsHandler) // GET /logs

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gateway is healthy!"))
	})

	return router
}
