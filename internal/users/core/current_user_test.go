package core

import (
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_GetReturnsSuccess_WhenDownstreamServicesAreOk(t *testing.T) {
	// Arrange
	fixture = newUsersServiceTestFixture()

	fixture.mockRepository.
		On("GetUserById", fixture.ctx, mock.AnythingOfType("uuid.UUID")).
		Return(&data.UserEntity{}, nil)

	fixture.mockTokenService.
		On("GenerateUserToken", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("string")).
		Return("stub token", nil)

	// Act
	response, err := fixture.service.Get(fixture.ctx, uuid.New())

	assert.Nil(t, err)
	assert.NotNil(t, response)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetUserById", 1)
	fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
