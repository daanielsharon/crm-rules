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

	action.CreatedAt = time.Now()
	action.UpdatedAt = time.Now()
	return s.Storage.CreateAction(*action)
}
