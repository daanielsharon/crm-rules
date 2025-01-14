package storage

import "log-worker/models"

type LogStorageInterface interface {
	CreateLog(log models.Log) error
}
