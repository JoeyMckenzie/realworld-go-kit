package specs

import (
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users/domain"
    "github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

func Test_WhenNoExistingUserIsFound_ReturnsSuccessfullyCreatedUser(t *testing.T) {
    // Arrange
    request := domain.RegisterUserServiceRequest{
        Email:    "test123@gmail.com",
        Username: "test123",
        Password: "test123",
    }

    fixture.mockSecurityService.
        On("HashPassword", mock.AnythingOfType("string")).
        Return("hashed password", nil)

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
        Return("stub token", nil)

    // Act
    response, err := fixture.service.RegisterUser(fixture.ctx, &request)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    assert.Equal(t, response.Username, request.Username)
    assert.Equal(t, response.Email, request.Email)
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
    fixture.resetMocks()
}

func Test_WhenExistingUsernameIsFound_ReturnsError(t *testing.T) {
    // Arrange
    request := domain.RegisterUserServiceRequest{
        Email:    "test123@gmail.com",
        Username: "testUser1",
        Password: "test123",
    }

    // Act
    response, err := fixture.service.RegisterUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrUsernameOrEmailTaken.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
    fixture.resetMocks()
}

func Test_WhenExistingEmailIsFound_ReturnsError(t *testing.T) {
    // Arrange
    request := domain.RegisterUserServiceRequest{
        Email:    "testUser1@gmail.com",
        Username: "test123",
        Password: "test123",
    }

    // Act
    response, err := fixture.service.RegisterUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrUsernameOrEmailTaken.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
    fixture.resetMocks()
}

func Test_WhenDownstreamsServicesFail_ReturnsError(t *testing.T) {
    // Arrange
    request := domain.RegisterUserServiceRequest{
        Email:    "test1234@gmail.com",
        Username: "test1234",
        Password: "test1234",
    }

    fixture.mockSecurityService.
        On("HashPassword", mock.AnythingOfType("string")).
        Return("", utilities.ErrMock)

    // Act
    response, err := fixture.service.RegisterUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrMock.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
    fixture.resetMocks()
}
