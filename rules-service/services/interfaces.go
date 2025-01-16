package services

import "rules-service/models"

type RuleServiceInterface interface {
	CreateRule(rule *models.Rule) error
	GetAllRules() ([]models.Rule, error)
	GetRule(id string) (*models.Rule, error)
	UpdateRule(rule *models.Rule) error
	DeleteRule(id string) error
}

type ActionServiceInterface interface {
	CreateAction(rule *models.Action) error
}
