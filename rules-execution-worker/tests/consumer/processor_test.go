package consumer_test

import (
	"errors"
	"task-execution-service/consumer"
	"task-execution-service/tests/mocks"
	"task-execution-service/types"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskProcessor_ProcessTask(t *testing.T) {
	t.Run("successful processing", func(t *testing.T) {
		db, sqlMock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockStorage := new(mocks.MockStorage)
		mockPublisher := new(mocks.MockPublisher)

		task := types.Task{
			RuleID:    123,
			Condition: "age > 25",
			Action:    "send_email",
		}

		mockRows := sqlMock.NewRows([]string{"id", "name", "email"}).
			AddRow("user-1", "John Doe", "john@example.com").
			AddRow("user-2", "Jane Doe", "jane@example.com")

		sqlMock.ExpectQuery("SELECT").WillReturnRows(mockRows)
		rows, err := db.Query("SELECT 1")
		assert.NoError(t, err)

		mockStorage.On("GetMatchingUsers", task.Condition).Return(rows, nil)
		mockPublisher.On("PublishLogs", mock.AnythingOfType("types.Log")).Return(nil).Times(2)

		processor := consumer.NewTaskProcessor(mockStorage, mockPublisher)
		err = processor.ProcessTask(task)

		assert.NoError(t, err)
		mockStorage.AssertExpectations(t)
		mockPublisher.AssertExpectations(t)
	})

	t.Run("storage error", func(t *testing.T) {
		db, sqlMock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockStorage := new(mocks.MockStorage)
		mockPublisher := new(mocks.MockPublisher)

		// Create test data
		task := types.Task{
			RuleID:    123,
			Condition: "age > 25",
			Action:    "send_email",
		}

		mockRows := sqlMock.NewRows([]string{"id", "name", "email"}).
			AddRow("user-1", "John Doe", "john@example.com")

		sqlMock.ExpectQuery("SELECT").WillReturnRows(mockRows)
		rows, err := db.Query("SELECT 1")
		assert.NoError(t, err)

		expectedErr := errors.New("database error")
		mockStorage.On("GetMatchingUsers", task.Condition).Return(rows, expectedErr)

		processor := consumer.NewTaskProcessor(mockStorage, mockPublisher)
		err = processor.ProcessTask(task)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "get matching users")
		mockStorage.AssertExpectations(t)
		mockPublisher.AssertNotCalled(t, "PublishLogs")
	})

	t.Run("publisher error", func(t *testing.T) {
		db, sqlMock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockStorage := new(mocks.MockStorage)
		mockPublisher := new(mocks.MockPublisher)

		task := types.Task{
			RuleID:    123,
			Condition: "age > 25",
			Action:    "send_email",
		}

		mockRows := sqlMock.NewRows([]string{"id", "name", "email"}).
			AddRow("user-1", "John Doe", "john@example.com")

		sqlMock.ExpectQuery("SELECT").WillReturnRows(mockRows)
		rows, err := db.Query("SELECT 1")
		assert.NoError(t, err)

		mockStorage.On("GetMatchingUsers", task.Condition).Return(rows, nil)
		mockPublisher.On("PublishLogs", mock.MatchedBy(func(log types.Log) bool {
			return true
		})).Return(errors.New("publisher error")).Times(1)

		processor := consumer.NewTaskProcessor(mockStorage, mockPublisher)
		err = processor.ProcessTask(task)

		assert.NoError(t, err) // process should continue even if publishing fails
		mockStorage.AssertExpectations(t)
		mockPublisher.AssertExpectations(t)
	})

	t.Run("row scan error", func(t *testing.T) {
		db, sqlMock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockStorage := new(mocks.MockStorage)
		mockPublisher := new(mocks.MockPublisher)

		task := types.Task{
			RuleID:    123,
			Condition: "age > 25",
			Action:    "send_email",
		}

		mockRows := sqlMock.NewRows([]string{"id", "name"}).
			AddRow("user-1", "John Doe")

		sqlMock.ExpectQuery("SELECT").WillReturnRows(mockRows)
		rows, err := db.Query("SELECT 1")
		assert.NoError(t, err)

		mockStorage.On("GetMatchingUsers", task.Condition).Return(rows, nil)

		processor := consumer.NewTaskProcessor(mockStorage, mockPublisher)
		err = processor.ProcessTask(task)

		assert.NoError(t, err) // process should continue despite scan errors
		mockStorage.AssertExpectations(t)
		mockPublisher.AssertNotCalled(t, "PublishLogs")
	})
}
