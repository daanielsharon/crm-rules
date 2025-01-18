package mocks

import (
	"rules-service/models"

	"github.com/stretchr/testify/mock"
)

type MockRuleStorage struct {
	mock.Mock
}

func (m *MockRuleStorage) CreateRule(rule models.Rule) error {
	args := m.Called(rule)
	return args.Error(0)
}

func (m *MockRuleStorage) GetAllRules() ([]models.Rule, error) {
	args := m.Called()
	return args.Get(0).([]models.Rule), args.Error(1)
}

func (m *MockRuleStorage) GetRuleById(id string) (*models.Rule, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Rule), args.Error(1)
}

func (m *MockRuleStorage) UpdateRule(rule models.Rule) error {
	args := m.Called(rule)
	return args.Error(0)
}

func (m *MockRuleStorage) DeleteRule(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
