package data

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

func (m *MockUsersRepository) GetUserById(ctx context.Context, id uuid.UUID) (*UserEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *MockUsersRepository) GetUserByUsernameAndEmail(ctx context.Context, username, email string) (*UserEntity, error) {
	args := m.Called(ctx, username, email)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *MockUsersRepository) GetUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *MockUsersRepository) SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error) {
	args := m.Called(ctx, username, email)
	return args.Get(0).([]UserEntity), args.Error(1)
}

func (m *MockUsersRepository) UpdateUser(ctx context.Context, id uuid.UUID, username, email, bio, image, password string) (*UserEntity, error) {
	args := m.Called(ctx, id, username, email, bio, image, password)
	return args.Get(0).(*UserEntity), args.Error(1)
}

func (m *MockUsersRepository) GetExistingFollow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) (uuid.UUID, error) {
	args := m.Called(ctx, followerId, followeeId)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *MockUsersRepository) AddFollow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) error {
	args := m.Called(ctx, followerId, followeeId)
	return args.Error(0)
}

func (m *MockUsersRepository) DeleteFollow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) error {
	args := m.Called(ctx, followerId, followeeId)
	return args.Error(1)
}
