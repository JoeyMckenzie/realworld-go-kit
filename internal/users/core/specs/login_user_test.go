package specs

import (
    "github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
    "github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

func Test_WhenUserExists_ReturnsProperlyMappedUser(t *testing.T) {
    // Arrange
    request := domain.LoginUserServiceRequest{
        Email:    "testUser1@gmail.com",
        Password: "testPassword1",
    }

    fixture.mockSecurityService.
        On("IsValidPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return(true)

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
        Return("stub token", nil)

    // Act
    response, err := fixture.service.LoginUser(fixture.ctx, &request)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    assert.Equal(t, response.Username, "testUser1")
    fixture.mockSecurityService.AssertNumberOfCalls(t, "IsValidPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
    fixture.resetMocks()
}

func Test_WhenUserDoesNotExist_ReturnsNotFound(t *testing.T) {
    // Arrange
    request := domain.LoginUserServiceRequest{
        Email:    "userDoesNotExist@gmail.com",
        Password: "testPassword1",
    }

    // Act
    response, err := fixture.service.LoginUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrUserNotFound.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "IsValidPassword", 0)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
    fixture.resetMocks()
}

func Test_WhenPasswordIsInvalid_ReturnsUnauthorized(t *testing.T) {
    // Arrange
    request := domain.LoginUserServiceRequest{
        Email:    "testUser1@gmail.com",
        Password: "badPassword1",
    }

    fixture.mockSecurityService.
        On("IsValidPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return(false)

    // Act
    response, err := fixture.service.LoginUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrInvalidLoginAttempt.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "IsValidPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
    fixture.resetMocks()
}

func Test_WhenDownstreamsFail_ReturnsError(t *testing.T) {
    // Arrange
    request := domain.LoginUserServiceRequest{
        Email:    "testUser1@gmail.com",
        Password: "testUser1",
    }

    fixture.mockSecurityService.
        On("IsValidPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return(true)

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
        Return("", utilities.ErrMock)

    // Act
    response, err := fixture.service.LoginUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrMock.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "IsValidPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
    fixture.resetMocks()
}
