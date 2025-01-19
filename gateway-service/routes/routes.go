package routes

import (
	"gateway/handlers"
	localMiddleware "gateway/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitializeRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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

	router.Route("/actions", func(r chi.Router) {
		r.Post("/", handlers.CreateActionHandler)       // POST /actions
		r.Put("/{id}", handlers.UpdateActionHandler)    // PUT /actions/:id
		r.Get("/{id}", handlers.GetActionHandler)       // GET /actions/:id
		r.Get("/", handlers.GetActionsHandler)          // GET /actions
		r.Delete("/{id}", handlers.DeleteActionHandler) // DELETE /actions/:id
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

	router.Route("/logs", func(r chi.Router) {
		r.Get("/", handlers.GetLogsHandler) // GET /logs
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetLogByIDHandler) // GET /logs/:id
		})
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gateway is healthy!"))
	})

	return router
}
