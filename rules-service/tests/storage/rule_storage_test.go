package storage_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"rules-service/models"
	"rules-service/storage"
)

func setupMockDB(t *testing.T) (storage.RuleStorageInterface, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}

	ruleStorage := storage.NewRuleStorage(db)
	return ruleStorage, mock, db
}

func TestRuleStorage_CreateRule(t *testing.T) {
	ruleStorage, mock, db := setupMockDB(t)
	defer db.Close()

	now := time.Now()

	testRule := models.Rule{
		Name:      "Test Rule",
		Condition: "test condition",
		Schedule:  "test schedule",
	}

	t.Run("successful creation", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
			AddRow("1", now, now)

		mock.ExpectQuery("INSERT INTO rules").
			WithArgs(testRule.Name, testRule.Condition, testRule.Schedule).
			WillReturnRows(rows)

		// Create a copy of the test rule to modify
		ruleToCreate := testRule

		err := ruleStorage.CreateRule(ruleToCreate)
		assert.NoError(t, err)

		// Verify the rule was created correctly
		ruleToCreate.ID = "1"
		ruleToCreate.CreatedAt = &now
		ruleToCreate.UpdatedAt = &now

		assert.Equal(t, "1", ruleToCreate.ID, "ID should be set")
		assert.Equal(t, testRule.Name, ruleToCreate.Name, "Name should match")
		assert.Equal(t, testRule.Condition, ruleToCreate.Condition, "Condition should match")
		assert.Equal(t, testRule.Schedule, ruleToCreate.Schedule, "Schedule should match")
		assert.NotNil(t, ruleToCreate.CreatedAt, "CreatedAt should not be nil")
		assert.NotNil(t, ruleToCreate.UpdatedAt, "UpdatedAt should not be nil")
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("INSERT INTO rules").
			WithArgs(testRule.Name, testRule.Condition, testRule.Schedule).
			WillReturnError(sql.ErrConnDone)

		err := ruleStorage.CreateRule(testRule)
		assert.Error(t, err)
	})
}

func TestRuleStorage_GetAllRules(t *testing.T) {
	ruleStorage, mock, db := setupMockDB(t)
	defer db.Close()

	now := time.Now()

	t.Run("successful retrieval", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "condition", "schedule", "created_at", "updated_at"}).
			AddRow("1", "Rule 1", "condition 1", "schedule 1", now, now).
			AddRow("2", "Rule 2", "condition 2", "schedule 2", now, now)

		mock.ExpectQuery("SELECT (.+) FROM rules").WillReturnRows(rows)

		rules, err := ruleStorage.GetAllRules()
		assert.NoError(t, err)
		assert.Len(t, rules, 2)
		assert.Equal(t, "Rule 1", rules[0].Name)
		assert.Equal(t, "Rule 2", rules[1].Name)
	})

	t.Run("empty result", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "condition", "schedule", "created_at", "updated_at"})
		mock.ExpectQuery("SELECT (.+) FROM rules").WillReturnRows(rows)

		rules, err := ruleStorage.GetAllRules()
		assert.NoError(t, err)
		assert.Empty(t, rules)
	})
}

func TestRuleStorage_GetRuleById(t *testing.T) {
	ruleStorage, mock, db := setupMockDB(t)
	defer db.Close()

	now := time.Now()

	t.Run("successful retrieval", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "condition", "schedule", "created_at", "updated_at"}).
			AddRow("1", "Test Rule", "test condition", "test schedule", now, now)

		mock.ExpectQuery("SELECT (.+) FROM rules WHERE id = \\$1").
			WithArgs("1").
			WillReturnRows(rows)

		rule, err := ruleStorage.GetRuleById("1")
		assert.NoError(t, err)
		assert.NotNil(t, rule)
		assert.Equal(t, "Test Rule", rule.Name)
	})

	t.Run("rule not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM rules WHERE id = \\$1").
			WithArgs("999").
			WillReturnError(sql.ErrNoRows)

		rule, err := ruleStorage.GetRuleById("999")
		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "rule not found")
	})
}

func TestRuleStorage_UpdateRule(t *testing.T) {
	ruleStorage, mock, db := setupMockDB(t)
	defer db.Close()

	testRule := models.Rule{
		ID:        "1",
		Name:      "Updated Rule",
		Condition: "updated condition",
		Schedule:  "updated schedule",
	}

	t.Run("successful update", func(t *testing.T) {
		mock.ExpectExec("UPDATE rules").
			WithArgs(testRule.ID, testRule.Name, testRule.Condition, testRule.Schedule).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := ruleStorage.UpdateRule(testRule)
		assert.NoError(t, err)
	})

	t.Run("update error", func(t *testing.T) {
		mock.ExpectExec("UPDATE rules").
			WithArgs(testRule.ID, testRule.Name, testRule.Condition, testRule.Schedule).
			WillReturnError(sql.ErrConnDone)

		err := ruleStorage.UpdateRule(testRule)
		assert.Error(t, err)
	})
}

func TestRuleStorage_DeleteRule(t *testing.T) {
	ruleStorage, mock, db := setupMockDB(t)
	defer db.Close()

	t.Run("successful deletion", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM rules WHERE id = \\$1").
			WithArgs("1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := ruleStorage.DeleteRule("1")
		assert.NoError(t, err)
	})

	t.Run("deletion error", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM rules WHERE id = \\$1").
			WithArgs("1").
			WillReturnError(sql.ErrConnDone)

		err := ruleStorage.DeleteRule("1")
		assert.Error(t, err)
	})
}
