package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/models"
	"user-service/services"
	"user-service/utils"
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
		utils.HandleRequestError(w, &utils.RequestError{
			Message: "Invalid request body",
			Status:  http.StatusBadRequest,
		})
		return
	}

	if err := h.Service.CreateUser(user); err != nil {
		utils.HandleRequestError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.HandleRequestError(w, utils.ErrMissingID)
		return
	}

	user, err := h.Service.GetUserById(id)
	if err != nil {
		utils.HandleRequestError(w, utils.ErrUserNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		utils.HandleRequestError(w, err)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.HandleRequestError(w, utils.ErrMissingID)
		return
	}

	_, err = h.Service.GetUserById(id)
	if err != nil {
		utils.HandleRequestError(w, utils.ErrUserNotFound)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleRequestError(w, &utils.RequestError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	user.ID = id
	if err := h.Service.UpdateUser(user); err != nil {
		utils.HandleRequestError(w, &utils.RequestError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	user.ID = ""
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.HandleRequestError(w, utils.ErrMissingID)
		return
	}

	_, err = h.Service.GetUserById(id)
	if err != nil {
		utils.HandleRequestError(w, utils.ErrUserNotFound)
		return
	}

	if err := h.Service.DeleteUser(id); err != nil {
		utils.HandleRequestError(w, &utils.RequestError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"user_id": id,
	})
}
