package handlers

import (
	"log-service/services"
	"log-service/utils"
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
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, logs, http.StatusOK)
}

func (h *LogHandler) GetLogByID(w http.ResponseWriter, r *http.Request) {
	logID := chi.URLParam(r, "id")
	if logID == "" {
		utils.ErrorResponse(w, "Log ID is required", http.StatusBadRequest)
		return
	}

	log, err := h.Service.GetLogByID(logID)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if log == nil {
		utils.ErrorResponse(w, "Log not found", http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, log, http.StatusOK)
}
