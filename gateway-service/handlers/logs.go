package handlers

import (
	"gateway/config"
	"gateway/utils"
	"io"
	"net/http"
)

func FetchLogsHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.RulesServiceURL, http.MethodGet, nil)
	if err != nil {
		http.Error(w, "Failed to fetch logs: "+err.Error(), http.StatusInternalServerError)
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
