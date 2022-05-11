package specs

import (
    "github.com/joeymckenzie/realworld-go-kit/conduit-domain/users"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

func Test_WhenUserExists_UpdatesFieldsAccordinglyAndReturnsMappedUser(t *testing.T) {
    // Arrange
    updatedImage := "testUser3 updated image"
    updatedBio := "testUser3 updated bio"
    request := users.UpdateUserServiceRequest{
        UserId: 3,
        Image:  &updatedImage,
        Bio:    &updatedBio,
    }

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
        Return("stub token", nil)

    // Act
    response, err := fixture.service.UpdateUser(fixture.ctx, &request)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    assert.Equal(t, response.Bio, *request.Bio)
    assert.Equal(t, response.Image, *request.Image)
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
    fixture.resetMocks()
}

func Test_WhenUserExistsWithUpdatedPassword_CallsSecurityService(t *testing.T) {
    // Arrange
    updatedPassword := "updated password"
    request := users.UpdateUserServiceRequest{
        UserId:   3,
        Password: &updatedPassword,
    }

    fixture.mockTokenService.
        On("GenerateUserToken", mock.AnythingOfType("int"), mock.AnythingOfType("string")).
        Return("updated hashed password", nil)

    fixture.mockSecurityService.
        On("HashPassword", mock.AnythingOfType("string")).
        Return("stub token", nil)

    // Act
    response, err := fixture.service.UpdateUser(fixture.ctx, &request)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 1)
    fixture.resetMocks()
}

func Test_WhenUserDoesNotExist_ReturnsWithUnauthorized(t *testing.T) {
    // Arrange
    request := users.UpdateUserServiceRequest{
        UserId: 33,
    }

    // Act
    response, err := fixture.service.UpdateUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrUnauthorized.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 0)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
    fixture.resetMocks()
}

func Test_WhenDownstreamServicesFail_ReturnsWithError(t *testing.T) {
    // Arrange
    updatedPassword := "updated password"
    request := users.UpdateUserServiceRequest{
        UserId:   3,
        Password: &updatedPassword,
    }

    fixture.mockSecurityService.
        On("HashPassword", mock.AnythingOfType("string")).
        Return("", utilities.ErrMock)

    // Act
    response, err := fixture.service.UpdateUser(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrMock.Error())
    fixture.mockSecurityService.AssertNumberOfCalls(t, "HashPassword", 1)
    fixture.mockTokenService.AssertNumberOfCalls(t, "GenerateUserToken", 0)
    fixture.resetMocks()
}
