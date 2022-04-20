package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_GetUser_WhenIsValidRequest_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return(stubToken, nil)

	// Act
	response, err := fixture.service.GetUser(fixture.ctx, 1)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}

func Test_GetUser_WhenUserRetrievalFails_ReturnsErr(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.GetUser(fixture.ctx, 1)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
}

func Test_GetUser_WhenTokenGenerationFails_ReturnsErr(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return("", utilities.ErrMock)

	// Act
	response, err := fixture.service.GetUser(fixture.ctx, 1)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
