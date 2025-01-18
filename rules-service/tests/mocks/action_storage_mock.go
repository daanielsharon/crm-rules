package mocks

import (
	"github.com/stretchr/testify/mock"

	"rules-service/models"
)

type MockActionStorage struct {
	mock.Mock
}

func (m *MockActionStorage) CreateAction(action models.Action) error {
	args := m.Called(action)
	return args.Error(0)
}

func (m *MockActionStorage) GetActions() ([]models.Action, error) {
	args := m.Called()
	return args.Get(0).([]models.Action), args.Error(1)
}

func (m *MockActionStorage) GetActionById(id string) (*models.Action, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Action), args.Error(1)
}

func (m *MockActionStorage) UpdateAction(action models.Action) error {
	args := m.Called(action)
	return args.Error(0)
}

func (m *MockActionStorage) DeleteAction(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
