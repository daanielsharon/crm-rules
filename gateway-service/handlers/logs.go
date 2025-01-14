package handlers

import (
	"gateway/config"
	"gateway/utils"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	serviceURLs := config.NewServiceURLs()
	response, err := utils.ForwardRequest(serviceURLs.LogServiceURL, http.MethodGet, nil)
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

func GetLogByIDHandler(w http.ResponseWriter, r *http.Request) {
	logID := chi.URLParam(r, "id")
	if logID == "" {
		http.Error(w, "Log ID is required", http.StatusBadRequest)
		return
	}

	serviceURLs := config.NewServiceURLs()
	url := serviceURLs.LogServiceURL + "/" + logID
	response, err := utils.ForwardRequest(url, http.MethodGet, nil)
	if err != nil {
		http.Error(w, "Failed to fetch log: "+err.Error(), http.StatusInternalServerError)
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
