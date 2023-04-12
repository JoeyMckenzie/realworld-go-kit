package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUsersRepository struct {
	mock.Mock
}

func (m *MockUsersRepository) ResetMocks() {
	m.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (m *MockUsersRepository) CreateUser(ctx context.Context, username, email, password string) (*UserEntity, error) {
	args := m.Called(ctx, username, email, password)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *MockUsersRepository) GetUser(ctx context.Context, id uuid.UUID) (*UserEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *MockUsersRepository) SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error) {
	args := m.Called(ctx, username, email)
	return args.Get(0).([]UserEntity), args.Error(1)
}
