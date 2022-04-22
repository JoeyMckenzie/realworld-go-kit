package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RemoveUserFollow_WhenUsersExists_ReturnsMappedUserFollow(t *testing.T) {
	// Arrange
	followerUserId := 3
	followeeUsername := "testUser2"

	// Act
	response, err := fixture.service.RemoveUserFollow(fixture.ctx, followerUserId, followeeUsername)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.False(t, response.Following)
}

func Test_RemoveUserFollow_WhenUserFolloweeDoesNotExist_ReturnsError(t *testing.T) {
	// Arrange
	followerUserId := 2
	followeeUsername := "userDoesNotExist"

	// Act
	response, err := fixture.service.RemoveUserFollow(fixture.ctx, followerUserId, followeeUsername)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrUserNotFound.Error())
}

func Test_WhenUserIsNotFollowingUserProfile_DoesNotReturnError(t *testing.T) {
	// Arrange
	followerUserId := 1
	followeeUsername := "testUser2"

	// Act
	response, err := fixture.service.RemoveUserFollow(fixture.ctx, followerUserId, followeeUsername)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.False(t, response.Following)
}
