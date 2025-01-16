package handlers

import (
	"encoding/json"
	"net/http"
	"shared/helpers"
	"user-service/models"
	"user-service/services"
)

type UserHandler struct {
	Service services.UserServiceInterface
}

func NewUserHandler(service services.UserServiceInterface) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateUser(user); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, user, http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromRequest(r)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUserById(id)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	helpers.JSONResponse(w, user, http.StatusOK)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, users, http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromRequest(r)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.Service.GetUserById(id)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id
	if err := h.Service.UpdateUser(user); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = ""
	helpers.JSONResponse(w, user, http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromRequest(r)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.Service.GetUserById(id)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.Service.DeleteUser(id); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, map[string]string{
		"user_id": id,
	}, http.StatusOK)
}
