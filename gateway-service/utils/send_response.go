package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(w http.ResponseWriter, response *http.Response) []byte {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return nil
	}

	var jsonData interface{}
	jsonErr := json.Unmarshal(body, &jsonData)

	standardResp := &StandardResponse{
		Code: response.StatusCode,
	}

	if jsonErr != nil {
		standardResp.Message = strings.TrimSpace(string(body))
		standardResp.Data = map[string]interface{}{}
	} else {
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			standardResp.Message = "success"
		} else {
			standardResp.Message = "failed"
		}

		standardResp.Data = jsonData
	}

	data, err := json.Marshal(standardResp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return nil
	}

	return data
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) []byte {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResp := &StandardResponse{
		Code:    statusCode,
		Message: "gateway error",
		Data: map[string]interface{}{
			"error": message,
		},
	}

	data, err := json.Marshal(errorResp)
	if err != nil {
		http.Error(w, "Failed to marshal error response", http.StatusInternalServerError)
		return nil
	}

	return data
}
