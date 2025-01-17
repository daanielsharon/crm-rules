package handlers

import (
	"encoding/json"
	"net/http"
	"rules-service/models"
	"rules-service/services"
	"shared/helpers"
)

type ActionHandler struct {
	Service services.ActionServiceInterface
}

func NewActionHandler(service services.ActionServiceInterface) *ActionHandler {
	return &ActionHandler{Service: service}
}

func (h *ActionHandler) CreateActionHandler(w http.ResponseWriter, r *http.Request) {
	var action models.Action

	if err := json.NewDecoder(r.Body).Decode(&action); err != nil {
		helpers.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateAction(&action); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.JSONResponse(w, action, http.StatusCreated)
}
