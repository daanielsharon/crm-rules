package handlers

import (
	"gateway/utils"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateRuleHandler(w http.ResponseWriter, r *http.Request) {
	response, err := utils.ForwardRequest("http://localhost:8081/rules", http.MethodPost, r.Body)
	if err != nil {
		http.Error(w, "Failed to create rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
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

	url := "http://localhost:8081/rules/" + ruleID
	response, err := utils.ForwardRequest(url, http.MethodPut, r.Body)
	if err != nil {
		http.Error(w, "Failed to update rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
