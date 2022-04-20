package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetUserProfile_WhenRequestIsValid_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", -1)

	// Assert
	assert.NotNil(t, response)
	assert.Equal(t, response.Following, false)
	assert.Nil(t, err)
}

func Test_GetUserProfile_WhenRequestIsValidAndUserIdIsProvided_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", 1)

	// Assert
	assert.NotNil(t, response)
	assert.Equal(t, response.Following, true)
	assert.Nil(t, err)
}

func Test_GetUserProfile_WhenRepositoryReturnsWithErrors_ReturnsErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", -1)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
}

func Test_GetUserProfile_WhenRepositoryNoUser_ReturnsErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", -1)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
}

func Test_GetUserProfile_WhenRepositoryUserFollowReturnsWithErrors_ReturnsErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", 2)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
}

func Test_GetUserProfile_WhenRepositoryUserFollowReturnsNoRows_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture(t)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", 2)

	// Assert
	assert.NotNil(t, response)
	assert.Equal(t, response.Following, false)
	assert.Nil(t, err)
}
