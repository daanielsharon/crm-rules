package services

import (
	"errors"
	"log-service/models"
	"log-service/storage"
)

type LogService struct {
	Storage storage.StorageInterface
}

func NewLogService(storage storage.StorageInterface) LogServiceInterface {
	return &LogService{Storage: storage}
}

func (ls *LogService) GetLogs(ruleID, userID string) ([]models.Log, error) {
	logs, err := ls.Storage.GetLogs(ruleID, userID)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (ls *LogService) GetLogByID(id string) (*models.Log, error) {
	log, err := ls.Storage.GetLogByID(id)
	if err != nil {
		return nil, err
	}
	if log == nil {
		return nil, errors.New("log not found")
	}
	return log, nil
}
