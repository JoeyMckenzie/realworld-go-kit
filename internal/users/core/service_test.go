package core

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_RegisterReturnsSuccess_WhenDownstreamServicesAreOk(t *testing.T) {
	// Arrange
	fixture = newUsersServiceTestFixture()
	stubEmail := "email@email.com"
	stubUsername := "username"
	stubPassword := "password"
	request := users.AuthenticationRequest[users.RegisterUserRequest]{
		User: &users.RegisterUserRequest{
			Email:    &stubEmail,
			Username: &stubUsername,
			Password: &stubPassword,
		},
	}

	fixture.mockRepository.
		On("SearchUsers", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return([]infrastructure.UserEntity{}, nil)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return("stub hashed password", nil)

	fixture.mockRepository.
		On("CreateUser", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&infrastructure.UserEntity{}, nil)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("string")).
		Return("stub token", nil)

	// Act
	response, err := fixture.service.Register(fixture.ctx, request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	fixture.mockRepository.AssertNumberOfCalls(t, "SearchUsers", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateUser", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
	fixture.resetMocks()
}

func Test_LoginReturnsSuccess_WhenDownstreamServicesAreOk(t *testing.T) {
	// Arrange
	fixture = newUsersServiceTestFixture()
	stubEmail := "email@email.com"
	stubPassword := "password"
	request := users.AuthenticationRequest[users.LoginUserRequest]{
		User: &users.LoginUserRequest{
			Email:    &stubEmail,
			Password: &stubPassword,
		},
	}

	fixture.mockRepository.
		On("SearchUsers", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return([]infrastructure.UserEntity{{}}, nil)

	fixture.mockSecurityService.
		On("IsValidPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("string")).
		Return("stub token", nil)

	// Act
	response, err := fixture.service.Login(fixture.ctx, request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	fixture.mockRepository.AssertNumberOfCalls(t, "SearchUsers", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "IsValidPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
	fixture.resetMocks()
}