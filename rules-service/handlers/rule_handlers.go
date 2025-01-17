package handlers

import (
	"encoding/json"
	"net/http"
	"rules-service/models"
	"rules-service/services"
	"shared/helpers"

	"github.com/go-chi/chi/v5"
)

type RuleHandler struct {
	Service services.RuleServiceInterface
}

func NewRuleHandler(service services.RuleServiceInterface) *RuleHandler {
	return &RuleHandler{Service: service}
}

func (h *RuleHandler) CreateRuleHandler(w http.ResponseWriter, r *http.Request) {
	var rule models.Rule

	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		helpers.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateRule(&rule); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.JSONResponse(w, rule, http.StatusCreated)
}

func (h *RuleHandler) GetAllRulesHandler(w http.ResponseWriter, r *http.Request) {
	rules, err := h.Service.GetAllRules()
	if err != nil {
		helpers.ErrorResponse(w, "Failed to fetch rules: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, rules, http.StatusOK)
}

func (h *RuleHandler) GetRuleById(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		helpers.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	rule, err := h.Service.GetRule(ruleID)
	if err != nil {
		helpers.ErrorResponse(w, "Failed to fetch rule", http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, rule, http.StatusOK)
}

func (h *RuleHandler) UpdateRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		helpers.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	var rule models.Rule
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		helpers.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	rule.ID = ruleID
	if err := h.Service.UpdateRule(&rule); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	rule.ID = ""
	helpers.JSONResponse(w, rule, http.StatusOK)
}

func (h *RuleHandler) DeleteRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "id")
	if ruleID == "" {
		helpers.ErrorResponse(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteRule(ruleID); err != nil {
		helpers.ErrorResponse(w, "Failed to delete rule", http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, []byte("Rule deleted successfully"), http.StatusOK)
}
