package routes

import (
	"rules-service/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRoutes(rules *handlers.RuleHandler, actions *handlers.ActionHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)    // Logs requests
	router.Use(middleware.Recoverer) // Recovers from panics

	router.Route("/rules", func(r chi.Router) {
		r.Post("/", rules.CreateRuleHandler)       // POST /rules
		r.Get("/", rules.GetAllRulesHandler)       // GET /rules
		r.Get("/{id}", rules.GetRuleById)          // GET /rules/{id}
		r.Put("/{id}", rules.UpdateRuleHandler)    // PUT /rules/{id}
		r.Delete("/{id}", rules.DeleteRuleHandler) // DELETE /rules/{id}

		r.Route("/actions", func(r chi.Router) {
			r.Post("/", actions.CreateActionHandler) // POST /rules/actions
		})

	})

	return router
}
