package handlers

import (
	"log-service/services"
	"net/http"
	"shared/helpers"

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
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, logs, http.StatusOK)
}

func (h *LogHandler) GetLogById(w http.ResponseWriter, r *http.Request) {
	logID := chi.URLParam(r, "id")
	if logID == "" {
		helpers.ErrorResponse(w, "Log ID is required", http.StatusBadRequest)
		return
	}

	log, err := h.Service.GetLogById(logID)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if log == nil {
		helpers.ErrorResponse(w, "Log not found", http.StatusNotFound)
		return
	}

	helpers.JSONResponse(w, log, http.StatusOK)
}
