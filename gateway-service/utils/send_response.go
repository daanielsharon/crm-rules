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
		standardResp.Message = "Success"
		standardResp.Data = jsonData
	}

	data, err := json.Marshal(standardResp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return nil
	}

	return data
}
