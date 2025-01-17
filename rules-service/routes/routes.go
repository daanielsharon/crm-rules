package routes

import (
	"rules-service/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRoutes(rules *handlers.RuleHandler, actions *handlers.ActionHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/rules", func(r chi.Router) {
		r.Get("/", rules.GetAllRulesHandler)
		r.Post("/", rules.CreateRuleHandler)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", rules.GetRuleById)
			r.Put("/", rules.UpdateRuleHandler)
			r.Delete("/", rules.DeleteRuleHandler)
		})
	})

	r.Route("/actions", func(r chi.Router) {
		r.Post("/", actions.CreateActionHandler)
	})

	return r
}
