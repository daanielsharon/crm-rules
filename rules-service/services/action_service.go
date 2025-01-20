package services

import (
	"errors"
	"rules-service/models"
	"rules-service/storage"
	"time"
)

type ActionService struct {
	Storage storage.ActionStorageInterface
}

func NewActionService(storage storage.ActionStorageInterface) *ActionService {
	return &ActionService{Storage: storage}
}

func (s *ActionService) CreateAction(action *models.Action) error {
	if action.RuleID == "" || action.Action == "" {
		return errors.New("all fields (rule_id, action) are required")
	}

	now := time.Now()
	action.CreatedAt = &now
	action.UpdatedAt = &now
	return s.Storage.CreateAction(*action)
}

func (s *ActionService) GetActions(ruleID string) ([]models.Action, error) {
	return s.Storage.GetActions(ruleID)
}

func (s *ActionService) GetActionById(id string) (*models.Action, error) {
	return s.Storage.GetActionById(id)
}

func (s *ActionService) UpdateAction(data *models.Action) error {
	if data.Action == "" {
		return errors.New("field action is required")
	}

	return s.Storage.UpdateAction(*data)
}

func (s *ActionService) DeleteAction(id string) error {
	return s.Storage.DeleteAction(id)
}
