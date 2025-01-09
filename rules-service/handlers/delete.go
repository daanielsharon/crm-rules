package handlers

import (
	"net/http"
	"rules/utils"

	"github.com/go-chi/chi/v5"
)

func (h *RuleHandler) DeleteRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		utils.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteRule(ruleID); err != nil {
		utils.ErrorResponse(w, "Failed to delete rule", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, []byte("Rule deleted successfully"), http.StatusOK)
}
