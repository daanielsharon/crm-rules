package handlers

import (
	"gateway/config"
	"gateway/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.UserServiceURL, http.MethodPost, r.Body)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		http.Error(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.UserServiceURL + ruleID
	response, err := utils.ForwardRequest(url, http.MethodPut, r.Body)
	if err != nil {
		http.Error(w, "Failed to update rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.UserServiceURL, http.MethodGet, nil)
	if err != nil {
		http.Error(w, "Failed to fetch all users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.UserServiceURL + id
	response, err := utils.ForwardRequest(url, http.MethodGet, nil)
	if err != nil {
		http.Error(w, "Failed to fetch user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.UserServiceURL + id
	response, err := utils.ForwardRequest(url, http.MethodDelete, nil)
	if err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}
