package handlers

import "rules/services"

type RuleHandler struct {
	Service services.RuleServiceInterface
}

func NewRuleHandler(service services.RuleServiceInterface) *RuleHandler {
	return &RuleHandler{Service: service}
}
