package storage

import "rules/models"

type RuleStorageInterface interface {
	CreateRule(rule models.Rule) error
	GetAllRules() ([]models.Rule, error)
	GetRule(id string) (*models.Rule, error)
	UpdateRule(rule models.Rule) error
	DeleteRule(id string) error
}
