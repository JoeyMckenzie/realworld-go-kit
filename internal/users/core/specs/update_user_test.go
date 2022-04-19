package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// newUsersServiceTestFixture sets up a common test fixture with in-place mocks for users service dependencies.
// Note that we don't need a validator dependency as validation is done within the service middleware.
func newUpdateUserTestFixture(t *testing.T) *usersServiceTestFixture {
	updatedEmail := "stub updated email"
	updatedUsername := "stub updated username"
	updatedPassword := "stub updated password"
	updatedBio := "stub updated bio"
	updatedImage := "stub updated image"
	stubUpdateUserRequest.Email = &updatedEmail
	stubUpdateUserRequest.Username = &updatedUsername
	stubUpdateUserRequest.Password = &updatedPassword
	stubUpdateUserRequest.Bio = &updatedBio
	stubUpdateUserRequest.Image = &updatedImage

	return newUsersServiceTestFixture(t)
}

func Test_UpdateUser_WhenRequestIsValidWithExistingUser_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return(stubToken, nil)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.IsType(t, &domain.UserDto{}, response)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}

func Test_UpdateUser_WhenRequestIsValidWithExistingUserAndNoPasswordUpdate_ReturnsSuccessfulResponseWithoutCallingHashPassword(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	stubUpdateUserRequest.Password = nil

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return(stubToken, nil)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.IsType(t, &domain.UserDto{}, response)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}

func Test_UpdateUser_WhenGetUserReturnsError_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Error(t, err, utilities.ErrUnauthorized)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_UpdateUser_WhenExistingUserIsNotFound_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Error(t, err, utilities.ErrUnauthorized)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_UpdateUser_WhenExistingUserIdDoesNotMatchRequestId_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Error(t, err, utilities.ErrUnauthorized)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_UpdateUser_WhenHashPasswordFails_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	fixture.mockSecurityService.
		On("HashPassword", mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_UpdateUser_WhenUpdateUserFails_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_UpdateUser_WhenGenerateTokenFails_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUpdateUserTestFixture(t)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.UpdateUser(fixture.ctx, &stubUpdateUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
