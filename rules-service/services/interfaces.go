package services

import "rules-service/models"

type RuleServiceInterface interface {
	CreateRule(rule *models.Rule) error
	GetAllRules() ([]models.Rule, error)
	GetRuleById(id string) (*models.Rule, error)
	UpdateRule(rule *models.Rule) error
	DeleteRule(id string) error
}

type ActionServiceInterface interface {
	CreateAction(rule *models.Action) error
	GetActions(rule_id string) ([]models.Action, error)
	GetActionById(id string) (*models.Action, error)
	UpdateAction(rule *models.Action) error
	DeleteAction(id string) error
}
