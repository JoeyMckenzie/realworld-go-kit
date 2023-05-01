package users

import (
    "github.com/gofrs/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

func Test_GetReturnsSuccess_WhenDownstreamServicesAreOk(t *testing.T) {
    // Arrange
    fixture = newUsersServiceTestFixture()

    fixture.mockRepository.
        On("GetUserById", fixture.ctx, mock.AnythingOfType("uuid.UUID")).
        Return(&repositories.UserEntity{}, nil)

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("string")).
        Return("stub token", nil)

    // Act
    id, _ := uuid.NewV4()
    response, err := fixture.service.Get(fixture.ctx, id)

    assert.Nil(t, err)
    assert.NotNil(t, response)
    fixture.mockRepository.AssertNumberOfCalls(t, "GetUserById", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
