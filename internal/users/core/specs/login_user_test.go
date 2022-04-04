package specs

import (
	"database/sql"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_LoginUser_WhenRequestIsValidWithNoExistingUser_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByEmail", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockSecurityService.
		On("PasswordIsValid", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return(stubToken, nil)

	// Act
	response, err := fixture.service.LoginUser(fixture.ctx, &stubLoginUserRequest)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.IsType(t, &domain.UserDto{}, response)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByEmail", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "PasswordIsValid", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}

func Test_LoginUser_WhenUserDoesNotExist_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByEmail", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, sql.ErrNoRows)

	// Act
	response, err := fixture.service.LoginUser(fixture.ctx, &stubLoginUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByEmail", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "PasswordIsValid", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_LoginUser_WhenRepositoryReturnsWithErrors_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByEmail", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.LoginUser(fixture.ctx, &stubLoginUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByEmail", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "PasswordIsValid", 0)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_LoginUser_WhenPasswordIsInvalid_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByEmail", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockSecurityService.
		On("PasswordIsValid", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(false)

	// Act
	response, err := fixture.service.LoginUser(fixture.ctx, &stubLoginUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByEmail", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "PasswordIsValid", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_LoginUser_WhenTokenServiceReturnsWithErrors_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByEmail", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockSecurityService.
		On("PasswordIsValid", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(true)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.LoginUser(fixture.ctx, &stubLoginUserRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByEmail", 1)
	fixture.mockSecurityService.AssertNumberOfCalls(t, "PasswordIsValid", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
