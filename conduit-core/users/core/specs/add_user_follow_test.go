package specs

import (
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_WhenUsersExists_ReturnsMappedUserFollowAsTrue(t *testing.T) {
    // Arrange
    followerUserId := 2
    followeeUsername := "testUser3"

    // Act
    response, err := fixture.service.AddUserFollow(fixture.ctx, followerUserId, followeeUsername)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    assert.True(t, response.Following)
}

func Test_WhenUserFolloweeDoesNotExist_ReturnsError(t *testing.T) {
    // Arrange
    followerUserId := 2
    followeeUsername := "userDoesNotExist"

    // Act
    response, err := fixture.service.AddUserFollow(fixture.ctx, followerUserId, followeeUsername)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrUserNotFound.Error())
}

func Test_WhenUserAttemptsToFollowSelf_ReturnsError(t *testing.T) {
    // Arrange
    followerUserId := 2
    followeeUsername := "testUser2"

    // Act
    response, err := fixture.service.AddUserFollow(fixture.ctx, followerUserId, followeeUsername)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrCannotFollowSelf.Error())
}
