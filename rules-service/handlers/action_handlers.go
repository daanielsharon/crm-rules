package handlers

import (
	"encoding/json"
	"net/http"
	"rules-service/models"
	"rules-service/services"
	"shared/helpers"

	"github.com/go-chi/chi/v5"
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

func (h *ActionHandler) GetActionsHandler(w http.ResponseWriter, r *http.Request) {
	actions, err := h.Service.GetActions()
	if err != nil {
		helpers.ErrorResponse(w, "Failed to fetch actions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, actions, http.StatusOK)
}

func (h *ActionHandler) GetActionById(w http.ResponseWriter, r *http.Request) {
	actionID := chi.URLParam(r, "id")
	if actionID == "" {
		helpers.ErrorResponse(w, "Action ID is required", http.StatusBadRequest)
		return
	}

	action, err := h.Service.GetActionById(actionID)
	if err != nil {
		helpers.ErrorResponse(w, "Failed to fetch action: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, action, http.StatusOK)
}

func (h *ActionHandler) UpdateActionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		helpers.ErrorResponse(w, "Action ID is required", http.StatusBadRequest)
		return
	}

	var action models.Action
	if err := json.NewDecoder(r.Body).Decode(&action); err != nil {
		helpers.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := h.Service.GetActionById(id)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	action.ID = id
	if err := h.Service.UpdateAction(&action); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	action.ID = ""
	helpers.JSONResponse(w, action, http.StatusOK)
}

func (h *ActionHandler) DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		helpers.ErrorResponse(w, "Action ID is required", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteAction(id); err != nil {
		helpers.ErrorResponse(w, "Failed to delete action: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, map[string]string{
		"action_id": id,
	}, http.StatusOK)
}
