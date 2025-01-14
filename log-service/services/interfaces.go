package services

import "log-service/models"

type LogServiceInterface interface {
	GetLogs(ruleID, userID string) ([]models.Log, error)
	GetLogByID(id string) (*models.Log, error)
}
