package handlers

import (
	"net/http"
	"rules/utils"

	"github.com/go-chi/chi/v5"
)

func (h *RuleHandler) GetAllRulesHandler(w http.ResponseWriter, r *http.Request) {
	rules, err := h.Service.GetAllRules()
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch rules", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, rules, http.StatusOK)
}

func (h *RuleHandler) GetRuleById(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		utils.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	rule, err := h.Service.GetRule(ruleID)
	if err != nil {
		utils.ErrorResponse(w, "Failed to fetch rule", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, rule, http.StatusOK)
}
