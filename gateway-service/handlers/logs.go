package handlers

import (
	"gateway/config"
	"gateway/utils"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()

	ruleID := r.URL.Query().Get("rule_id")
	userID := r.URL.Query().Get("user_id")

	fullURL := strings.TrimSuffix(serviceURLs.LogServiceURL, "/")

	var queryParams []string
	if ruleID != "" {
		queryParams = append(queryParams, "rule_id="+ruleID)
	}
	if userID != "" {
		queryParams = append(queryParams, "user_id="+userID)
	}

	joinedParams := strings.TrimSuffix(strings.Join(queryParams, "&"), "&")
	fullURL += "?" + joinedParams

	response, err := utils.ForwardRequest(fullURL, http.MethodGet, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch logs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}

func GetLogByIDHandler(w http.ResponseWriter, r *http.Request) {
	logID := chi.URLParam(r, "id")
	if logID == "" {
		utils.ErrorResponse(w, "Log ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.LogServiceURL + logID
	response, err := utils.ForwardRequest(url, http.MethodGet, nil)
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch log: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(utils.SendResponse(w, response))
}
