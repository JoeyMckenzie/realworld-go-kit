package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type mockUsersRepository struct {
	mock.Mock
}

func (m *mockUsersRepository) ResetMocks() {
	m.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (m *mockUsersRepository) CreateUser(ctx context.Context, username, email, password string) (*UserEntity, error) {
	args := m.Called(ctx, username, email, password)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *mockUsersRepository) GetUser(ctx context.Context, id uuid.UUID) (*UserEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *mockUsersRepository) SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error) {
	args := m.Called(ctx, username, email)
	return args.Get(0).([]UserEntity), args.Error(1)
}
