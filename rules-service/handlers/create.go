package handlers

import (
	"encoding/json"
	"net/http"
	"rules/models"
	"rules/utils"
)

func (h *RuleHandler) CreateRuleHandler(w http.ResponseWriter, r *http.Request) {
	var rule models.Rule

	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateRule(&rule); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, rule, http.StatusCreated)
}
