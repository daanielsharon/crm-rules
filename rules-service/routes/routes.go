package routes

import (
	"rules/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRoutes(handler *handlers.RuleHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)    // Logs requests
	router.Use(middleware.Recoverer) // Recovers from panics

	router.Route("/rules", func(r chi.Router) {
		r.Post("/", handler.CreateRuleHandler)       // POST /rules
		r.Get("/", handler.GetAllRulesHandler)       // GET /rules
		r.Get("/{id}", handler.GetRuleById)          // GET /rules/{id}
		r.Put("/{id}", handler.UpdateRuleHandler)    // PUT /rules/{id}
		r.Delete("/{id}", handler.DeleteRuleHandler) // DELETE /rules/{id}
	})

	return router
}
