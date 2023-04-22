package users

import (
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

func Test_UpdateReturnsSuccess_WhenDownstreamServicesAreOk(t *testing.T) {
    // Arrange
    fixture = newUsersServiceTestFixture()
    request := domain.AuthenticationRequest[domain.UpdateUserRequest]{
        User: &domain.UpdateUserRequest{
            Email:    stubEmail,
            Password: stubPassword,
            Image:    stubImage,
            Bio:      stubBio,
            Username: stubUsername,
        },
    }

    fixture.mockRepository.
        On("GetUserById", fixture.ctx, mock.AnythingOfType("uuid.UUID")).
        Return(&repositories.UserEntity{}, nil)

    fixture.mockRepository.
        On("SearchUsers", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return([]repositories.UserEntity{}, nil)

    fixture.mockSecurityService.
        On("HashPassword", mock.AnythingOfType("string")).
        Return(stubPassword, nil)

    fixture.mockRepository.
        On("UpdateUser",
            fixture.ctx,
            mock.AnythingOfType("uuid.UUID"),
            mock.AnythingOfType("string"),
            mock.AnythingOfType("string"),
            mock.AnythingOfType("string"),
            mock.AnythingOfType("string"),
            mock.AnythingOfType("string")).
        Return(&repositories.UserEntity{}, nil)

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("string")).
        Return("stub token", nil)

    // Act
    response, err := fixture.service.Update(fixture.ctx, request, uuid.New())

    assert.Nil(t, err)
    assert.NotNil(t, response)
    fixture.mockRepository.AssertNumberOfCalls(t, "GetUserById", 1)
    fixture.mockRepository.AssertNumberOfCalls(t, "SearchUsers", 1)
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
    fixture.mockRepository.AssertNumberOfCalls(t, "UpdateUser", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
