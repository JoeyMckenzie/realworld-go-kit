package users

import (
    "github.com/go-faker/faker/v4"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

func Test_LoginReturnsSuccess_WhenDownstreamServicesAreOk(t *testing.T) {
    // Arrange
    fixture = newUsersServiceTestFixture()
    request := domain.AuthenticationRequest[domain.LoginUserRequest]{
        User: &domain.LoginUserRequest{
            Email:    faker.Email(),
            Password: faker.Password(),
        },
    }

    fixture.mockRepository.
        On("SearchUsers", fixture.ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return([]repositories.UserEntity{{}}, nil)

    fixture.mockSecurityService.
        On("IsValidPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return(true)

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("string")).
        Return("stub token", nil)

    // Act
    response, err := fixture.service.Login(fixture.ctx, request)

    assert.Nil(t, err)
    assert.NotNil(t, response)
    fixture.mockRepository.AssertNumberOfCalls(t, "SearchUsers", 1)
    fixture.mockSecurityService.AssertNumberOfCalls(t, "IsValidPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
}
