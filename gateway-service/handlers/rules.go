package handlers

import (
	"gateway/config"
	"gateway/utils"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateRuleHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.RulesServiceURL, http.MethodPost, r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to create rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.ErrorResponse(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	w.Write(body)
}

func UpdateRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		http.Error(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.RulesServiceURL + "/" + ruleID
	response, err := utils.ForwardRequest(url, http.MethodPut, r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to update rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, response)
}

func GetRulesHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.RulesServiceURL, http.MethodGet, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch rules: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func GetRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		utils.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.RulesServiceURL + "/" + ruleID
	response, err := utils.ForwardRequest(url, http.MethodGet, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func DeleteRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		utils.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.RulesServiceURL + "/" + ruleID
	response, err := utils.ForwardRequest(url, http.MethodDelete, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to delete rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}
