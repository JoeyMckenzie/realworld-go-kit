package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_RegisterUser_WhenRequestIsValidWithNoExistingUser_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)
	defer fixture.client.Close()

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return(stubToken, nil)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return("stub hashed password", nil)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.IsType(t, &domain.UserDto{}, response)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}

func Test_RegisterUser_WhenUserExists_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenUserRepositoryReturnsError_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenHashingServiceReturnsWithError_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenCreateUserRepositoryReturnsWithErrors_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return(stubRegisterUserRequest.Password, nil)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_RegisterUser_WhenTokenServiceReturnsWithErrors_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return(stubRegisterUserRequest.Password, nil)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.RegisterUser(fixture.ctx, &stubRegisterUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
