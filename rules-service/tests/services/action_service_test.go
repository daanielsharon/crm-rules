package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"rules-service/models"
	"rules-service/services"
	"rules-service/tests/mocks"
)

func TestActionService_CreateAction(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		mockStorage := new(mocks.MockActionStorage)
		actionService := services.NewActionService(mockStorage)

		testAction := &models.Action{
			RuleID: "rule-123",
			Action: "test action",
		}

		mockStorage.On("CreateAction", mock.Anything).Return(nil)

		err := actionService.CreateAction(testAction)
		assert.NoError(t, err)
		assert.NotNil(t, testAction.CreatedAt)
		assert.NotNil(t, testAction.UpdatedAt)
		assert.Equal(t, testAction.RuleID, "rule-123")
		assert.Equal(t, testAction.Action, "test action")

		mockStorage.AssertExpectations(t)
	})

	t.Run("missing rule id", func(t *testing.T) {
		mockStorage := new(mocks.MockActionStorage)
		actionService := services.NewActionService(mockStorage)

		testAction := &models.Action{
			Action: "test action",
		}

		err := actionService.CreateAction(testAction)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "rule_id")

		mockStorage.AssertNotCalled(t, "CreateAction")
	})

	t.Run("missing action", func(t *testing.T) {
		mockStorage := new(mocks.MockActionStorage)
		actionService := services.NewActionService(mockStorage)

		testAction := &models.Action{
			RuleID: "rule-123",
		}

		err := actionService.CreateAction(testAction)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "action")
		mockStorage.AssertNotCalled(t, "CreateAction")
	})

	t.Run("storage error", func(t *testing.T) {
		mockStorage := new(mocks.MockActionStorage)

		actionService := services.NewActionService(mockStorage)

		testAction := &models.Action{
			RuleID: "rule-123",
			Action: "test action",
		}

		mockStorage.On("CreateAction", mock.AnythingOfType("models.Action")).Return(errors.New("storage error"))

		err := actionService.CreateAction(testAction)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "storage error")

		mockStorage.AssertExpectations(t)
	})
}
