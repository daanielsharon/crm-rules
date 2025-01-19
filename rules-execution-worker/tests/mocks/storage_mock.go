package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetMatchingUsers(condition string) (*sql.Rows, error) {
	args := m.Called(condition)
	return args.Get(0).(*sql.Rows), args.Error(1)
}
