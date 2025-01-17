package app

import (
	"rules-service/handlers"
	"rules-service/routes"
	"rules-service/services"
	"rules-service/storage"
	"shared/db"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	database := db.InitPostgres()
	ruleStorage := storage.NewRuleStorage(database)
	ruleService := services.NewRuleService(ruleStorage)
	rulesHandler := handlers.NewRuleHandler(ruleService)

	actionStorage := storage.NewActionStorage(database)
	actionService := services.NewActionService(actionStorage)
	actionHandler := handlers.NewActionHandler(actionService)

	return routes.InitializeRoutes(rulesHandler, actionHandler)
}
