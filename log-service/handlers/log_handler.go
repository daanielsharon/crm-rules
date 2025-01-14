package handlers

import (
	"encoding/json"
	"log-service/services"
	"net/http"

	"github.com/go-chi/chi"
)

type LogHandler struct {
	Service services.LogServiceInterface
}

func NewLogHandler(service services.LogServiceInterface) *LogHandler {
	return &LogHandler{Service: service}
}

func (h *LogHandler) GetLogs(w http.ResponseWriter, r *http.Request) {
	ruleID := r.URL.Query().Get("rule_id")
	userID := r.URL.Query().Get("user_id")

	logs, err := h.Service.GetLogs(ruleID, userID)
	if err != nil {
		http.Error(w, "Failed to retrieve logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func (h *LogHandler) GetLogByID(w http.ResponseWriter, r *http.Request) {
	logID := chi.URLParam(r, "id")
	if logID == "" {
		http.Error(w, "Log ID is required", http.StatusBadRequest)
		return
	}

	log, err := h.Service.GetLogByID(logID)
	if err != nil {
		http.Error(w, "Failed to retrieve log", http.StatusInternalServerError)
		return
	}
	if log == nil {
		http.Error(w, "Log not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(log)
}
