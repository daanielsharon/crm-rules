package storage

import "rules-service/models"

type RuleStorageInterface interface {
	CreateRule(rule models.Rule) error
	GetAllRules() ([]models.Rule, error)
	GetRuleById(id string) (*models.Rule, error)
	UpdateRule(rule models.Rule) error
	DeleteRule(id string) error
}

type ActionStorageInterface interface {
	CreateAction(action models.Action) error
	GetActions(ruleID string) ([]models.Action, error)
	GetActionById(id string) (*models.Action, error)
	UpdateAction(action models.Action) error
	DeleteAction(id string) error
}
