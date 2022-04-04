package persistence

import (
	"context"
	"github.com/stretchr/testify/mock"
	"time"
)

var (
	MockUser = &UserEntity{
		Id:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  "stub password",
		Email:     "stub email",
		Username:  "stub username",
		Bio:       "stub bio",
		Image:     "stub image",
	}

	AnotherMockUser = &UserEntity{
		Id:        2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  "another stub password",
		Email:     "another stub email",
		Username:  "another stub username",
		Bio:       "another stub bio",
		Image:     "another stub image",
	}

	MockUserProfileFollow = &UserProfileFollowEntity{
		Id:             1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		FolloweeUserId: 1,
		FollowerUserId: 2,
	}
)

type MockUsersRepository struct {
	mock.Mock
}

func NewMockUsersRepository() UsersRepository {
	return &MockUsersRepository{}
}

func (m *MockUsersRepository) GetUser(ctx context.Context, userId int) (*UserEntity, error) {
	args := m.Called(ctx, userId)
	return handleNilUserMockOrDefault[UserEntity](args)
}

func (m *MockUsersRepository) CreateUser(ctx context.Context, username, email, password string) (*UserEntity, error) {
	args := m.Called(ctx, username, email, password)
	return handleNilUserMockOrDefault[UserEntity](args)
}

func (m *MockUsersRepository) UpdateUser(ctx context.Context, userId int, username, email, password, bio, image string) (*UserEntity, error) {
	args := m.Called(ctx, userId, username, email, password, image, bio)
	return handleNilUserMockOrDefault[UserEntity](args)
}

func (m *MockUsersRepository) FindUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
	args := m.Called(ctx, username)
	return handleNilUserMockOrDefault[UserEntity](args)
}

func (m *MockUsersRepository) FindUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
	args := m.Called(ctx, email)
	return handleNilUserMockOrDefault[UserEntity](args)
}

func (m *MockUsersRepository) FindUserByUsernameOrEmail(ctx context.Context, username, email string) (*UserEntity, error) {
	args := m.Called(ctx, username, email)
	return handleNilUserMockOrDefault[UserEntity](args)
}

func (m *MockUsersRepository) GetUserProfileFollowByFollowee(ctx context.Context, followerUserId, followeeUserId int) (*UserProfileFollowEntity, error) {
	args := m.Called(ctx, followerUserId, followeeUserId)
	return handleNilUserMockOrDefault[UserProfileFollowEntity](args)
}

func (m *MockUsersRepository) CreateUserFollow(ctx context.Context, followerUserId, followeeUserId int) (*UserProfileFollowEntity, error) {
	args := m.Called(ctx, followerUserId, followeeUserId)
	return handleNilUserMockOrDefault[UserProfileFollowEntity](args)
}

func (m *MockUsersRepository) RemoveUserFollow(ctx context.Context, followerUserId, followeeUserId int) error {
	args := m.Called(ctx, followerUserId, followeeUserId)
	return args.Error(0)
}

func handleNilUserMockOrDefault[T UserProfileFollowEntity | UserEntity](args mock.Arguments) (*T, error) {
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*T), args.Error(1)
}
