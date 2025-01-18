package mocks

import (
	"task-execution-service/types"

	"github.com/stretchr/testify/mock"
)

type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) PublishLogs(log types.Log) error {
	args := m.Called(log)
	return args.Error(0)
}
