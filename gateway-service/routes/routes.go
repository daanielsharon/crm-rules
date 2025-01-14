package routes

import (
	"gateway/handlers"
	localMiddleware "gateway/middleware"
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
		r.Post("/", handlers.CreateRuleHandler)       // POST /rules
		r.Put("/{id}", handlers.UpdateRuleHandler)    // PUT /rules/:id
		r.Get("/{id}", handlers.GetRuleHandler)       // GET /rules/:id
		r.Get("/", handlers.GetRulesHandler)          // GET /rules
		r.Delete("/{id}", handlers.DeleteRuleHandler) // DELETE /rules/:id
	})

	router.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.CreateUserHandler) // POST /users
		r.Get("/", handlers.GetAllUsersHandler) // GET /users
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetUserHandler)       // GET /users/:id
			r.Put("/", handlers.UpdateUserHandler)    // PUT /users/:id
			r.Delete("/", handlers.DeleteUserHandler) // DELETE /users/:id
		})
	})

	router.Get("/logs", handlers.GetLogsHandler) // GET /logs
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gateway is healthy!"))
	})

	return router
}
