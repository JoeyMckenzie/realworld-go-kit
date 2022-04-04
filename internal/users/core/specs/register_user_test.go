package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_RegisterUser_WhenRequestIsValidWithNoExistingUser_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsernameOrEmail", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return(stubRegisterUserRequest.Password, nil)

	fixture.mockRepository.
		On("CreateUser", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return(stubToken, nil)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.IsType(t, &domain.UserDto{}, response)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsernameOrEmail", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateUser", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}

func Test_RegisterUser_WhenUserExists_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsernameOrEmail", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsernameOrEmail", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateUser", 0)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenUserRepositoryReturnsError_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsernameOrEmail", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsernameOrEmail", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateUser", 0)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenHashingServiceReturnsWithError_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsernameOrEmail", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsernameOrEmail", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateUser", 0)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenCreateUserRepositoryReturnsWithErrors_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsernameOrEmail", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return(stubRegisterUserRequest.Password, nil)

	fixture.mockRepository.
		On("CreateUser", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsernameOrEmail", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateUser", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenTokenServiceReturnsWithErrors_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsernameOrEmail", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return(stubRegisterUserRequest.Password, nil)

	fixture.mockRepository.
		On("CreateUser", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsernameOrEmail", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateUser", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
