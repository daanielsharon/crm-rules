package storage

import "rules-service/models"

type RuleStorageInterface interface {
	CreateRule(rule models.Rule) error
	GetAllRules() ([]models.Rule, error)
	GetRule(id string) (*models.Rule, error)
	UpdateRule(rule models.Rule) error
	DeleteRule(id string) error
}

type ActionStorageInterface interface {
	CreateAction(rule models.Action) error
	GetActions() ([]models.Action, error)
	GetActionById(id string) (*models.Action, error)
	UpdateAction(rule models.Action) error
	DeleteAction(id string) error
}
