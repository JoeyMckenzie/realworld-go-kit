package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_WhenLoginUserExists_ReturnsProperlyMappedUser(t *testing.T) {
	// Arrange
	userId := 1

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return("stub token", nil)

	// Act
	response, err := fixture.service.GetUser(fixture.ctx, userId)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, response.Username, "testUser1")
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
	fixture.resetMocks()
}

func Test_WhenUserDoesNotExist_ReturnsError(t *testing.T) {
	// Arrange
	userId := 11

	// Act
	response, err := fixture.service.GetUser(fixture.ctx, userId)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrUserNotFound.Error())
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
	fixture.resetMocks()
}

func Test_WhenServicesFail_ReturnsError(t *testing.T) {
	// Arrange
	userId := 1

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.GetUser(fixture.ctx, userId)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrMock.Error())
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
	fixture.resetMocks()
}
