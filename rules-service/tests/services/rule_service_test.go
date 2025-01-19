package services_test

import (
	"rules-service/models"
	"rules-service/services"
	"rules-service/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRule(t *testing.T) {
	t.Run("successful rule creation", func(t *testing.T) {
		mockStorage := new(mocks.MockRuleStorage)
		ruleService := services.NewRuleService(mockStorage)

		rule := &models.Rule{
			Name:      "Test Rule",
			Condition: "user.age > 25",
			Schedule:  "0 * * * *",
		}

		mockStorage.On("CreateRule", mock.AnythingOfType("models.Rule")).Return(nil)

		err := ruleService.CreateRule(rule)
		assert.NoError(t, err)
		assert.NotNil(t, rule.CreatedAt)
		assert.NotNil(t, rule.UpdatedAt)
		mockStorage.AssertExpectations(t)
	})

	t.Run("missing required fields", func(t *testing.T) {
		mockStorage := new(mocks.MockRuleStorage)
		ruleService := services.NewRuleService(mockStorage)

		testCases := []struct {
			name     string
			rule     *models.Rule
			errorMsg string
		}{
			{
				name: "missing name",
				rule: &models.Rule{
					Condition: "user.age > 25",
					Schedule:  "0 * * * *",
				},
				errorMsg: "all fields (name, condition, schedule) are required",
			},
			{
				name: "missing condition",
				rule: &models.Rule{
					Name:     "Test Rule",
					Schedule: "0 * * * *",
				},
				errorMsg: "all fields (name, condition, schedule) are required",
			},
			{
				name: "missing schedule",
				rule: &models.Rule{
					Name:      "Test Rule",
					Condition: "user.age > 25",
				},
				errorMsg: "all fields (name, condition, schedule) are required",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				err := ruleService.CreateRule(tc.rule)
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
			})
		}
	})

	t.Run("storage error", func(t *testing.T) {
		mockStorage := new(mocks.MockRuleStorage)
		ruleService := services.NewRuleService(mockStorage)

		rule := &models.Rule{
			Name:      "Test Rule",
			Condition: "user.age > 25",
			Schedule:  "0 * * * *",
		}

		expectedErr := assert.AnError
		mockStorage.On("CreateRule", mock.AnythingOfType("models.Rule")).Return(expectedErr)

		err := ruleService.CreateRule(rule)
		assert.ErrorIs(t, err, expectedErr)
		mockStorage.AssertExpectations(t)
	})
}
