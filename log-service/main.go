package main

import (
	"log"
	"log-service/db"
	"log-service/handlers"
	"log-service/services"
	"log-service/storage"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	store := storage.NewStorage(db.InitPostgres())
	service := services.NewLogService(store)
	handler := handlers.NewLogHandler(service)

	r := chi.NewRouter()
	r.Get("/logs", handler.GetLogs)

	log.Println("Log service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
