package specs

import (
	"database/sql"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_GetUserProfile_WhenRequestIsValid_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsername", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", -1)

	// Assert
	assert.NotNil(t, response)
	assert.Equal(t, response.Following, false)
	assert.Nil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsername", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetUserProfileFollowByFollowee", 0)
}

func Test_GetUserProfile_WhenRequestIsValidAndUserIdIsProvided_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsername", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockRepository.
		On("GetUserProfileFollowByFollowee", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(persistence.MockUserProfileFollow, nil)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", 1)

	// Assert
	assert.NotNil(t, response)
	assert.Equal(t, response.Following, true)
	assert.Nil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsername", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetUserProfileFollowByFollowee", 1)
}

func Test_GetUserProfile_WhenRepositoryReturnsWithErrors_ReturnsErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsername", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", -1)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsername", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetUserProfileFollowByFollowee", 0)
}

func Test_GetUserProfile_WhenRepositoryNoUser_ReturnsErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsername", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, sql.ErrNoRows)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", -1)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsername", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetUserProfileFollowByFollowee", 0)
}

func Test_GetUserProfile_WhenRepositoryUserFollowReturnsWithErrors_ReturnsErrors(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsername", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockRepository.
		On("GetUserProfileFollowByFollowee", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", 2)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsername", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetUserProfileFollowByFollowee", 1)
}

func Test_GetUserProfile_WhenRepositoryUserFollowReturnsNoRows_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newUsersServiceTestFixture()

	fixture.mockRepository.
		On("FindUserByUsername", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.MockUser, nil)

	fixture.mockRepository.
		On("GetUserProfileFollowByFollowee", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(nil, sql.ErrNoRows)

	// Act
	response, err := fixture.service.GetUserProfile(fixture.ctx, "userOne", 2)

	// Assert
	assert.NotNil(t, response)
	assert.Equal(t, response.Following, false)
	assert.Nil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindUserByUsername", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetUserProfileFollowByFollowee", 1)
}
