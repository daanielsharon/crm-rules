package handlers

import (
	"gateway/config"
	"gateway/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateActionHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.RulesServiceURL, http.MethodPost, r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to create actions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func UpdateActionHandler(w http.ResponseWriter, r *http.Request) {
	actionID := chi.URLParam(r, "id")
	if actionID == "" {
		utils.ErrorResponse(w, "Action ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.RulesServiceURL + actionID
	response, err := utils.ForwardRequest(url, http.MethodPut, r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to update rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func GetActionsHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.RulesServiceURL, http.MethodGet, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch actions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func GetActionHandler(w http.ResponseWriter, r *http.Request) {
	actionID := chi.URLParam(r, "id")
	if actionID == "" {
		utils.ErrorResponse(w, "Action ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.RulesServiceURL + actionID
	response, err := utils.ForwardRequest(url, http.MethodGet, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch action: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	actionID := chi.URLParam(r, "id")
	if actionID == "" {
		utils.ErrorResponse(w, "Action ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.RulesServiceURL + actionID
	response, err := utils.ForwardRequest(url, http.MethodDelete, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to delete action: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}
