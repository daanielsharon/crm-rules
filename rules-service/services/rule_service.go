package services

import (
	"errors"
	"rules-service/models"
	"rules-service/storage"
	"time"
)

type RuleService struct {
	Storage storage.RuleStorageInterface
}

func NewRuleService(storage storage.RuleStorageInterface) *RuleService {
	return &RuleService{Storage: storage}
}

func (s *RuleService) CreateRule(rule *models.Rule) error {
	if rule.Name == "" || rule.Condition == "" || rule.Schedule == "" {
		return errors.New("all fields (name, condition, schedule) are required")
	}

	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()
	return s.Storage.CreateRule(*rule)
}

func (s *RuleService) GetAllRules() ([]models.Rule, error) {
	return s.Storage.GetAllRules()
}

func (s *RuleService) GetRuleById(id string) (*models.Rule, error) {
	return s.Storage.GetRuleById(id)
}

func (s *RuleService) UpdateRule(rule *models.Rule) error {
	if rule.Name == "" || rule.Condition == "" || rule.Schedule == "" {
		return errors.New("all fields (name, condition, schedule) are required")
	}

	rule.UpdatedAt = time.Now()
	return s.Storage.UpdateRule(*rule)
}

func (s *RuleService) DeleteRule(id string) error {
	return s.Storage.DeleteRule(id)
}
