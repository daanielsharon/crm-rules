package storage

import "log-service/models"

type StorageInterface interface {
	GetLogs(ruleID, userID string) ([]models.Log, error)
	GetLogByID(id string) (*models.Log, error)
}
