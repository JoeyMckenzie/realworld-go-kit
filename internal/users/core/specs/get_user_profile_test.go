package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenUserExistsAndNotFollowingProfile_ReturnsFollowingFalse(t *testing.T) {
	// Arrange
	username := "testUser2"
	currentUserId := 1

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, username, currentUserId)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.False(t, response.Following)
}

func Test_WhenUserExistsAndIsFollowingProfile_ReturnsFollowingTrue(t *testing.T) {
	// Arrange
	username := "testUser1"
	currentUserId := 2

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, username, currentUserId)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.True(t, response.Following)
}

func Test_WhenUserProfileDoesNotExist_ReturnsError(t *testing.T) {
	// Arrange
	username := "userThatDoesNotExist"
	currentUserId := 2

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, username, currentUserId)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrUserNotFound.Error())
}
