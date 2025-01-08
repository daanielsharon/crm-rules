package handlers

import (
	"encoding/json"
	"net/http"
	"rules/models"
	"rules/utils"

	"github.com/go-chi/chi/v5"
)

func (h *RuleHandler) UpdateRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		utils.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	var rule models.Rule
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	rule.ID = ruleID
	if err := h.Service.UpdateRule(&rule); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, rule, http.StatusOK)
}
