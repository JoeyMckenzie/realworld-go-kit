package specs

import (
	"context"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
)

var (
	stubToken               = "stub token"
	stubRegisterUserRequest = domain.RegisterUserServiceRequest{
		Email:    "stub email",
		Username: "stub username",
		Password: "stub password",
	}
	stubLoginUserRequest = domain.LoginUserServiceRequest{
		Email:    "stub email",
		Password: "stub password",
	}
	stubUpdateUserRequest = domain.UpdateUserServiceRequest{
		UserId: 1,
	}
)

type usersServiceTestFixture struct {
	mockTokenService    *services.MockTokenService
	mockSecurityService *services.MockSecurityService
	mockRepository      *persistence.MockUsersRepository
	service             core.UsersService
	ctx                 context.Context
}

// newUsersServiceTestFixture sets up a common test fixture with in-place mocks for users service dependencies.
// Note that we don't need a validator dependency as validation is done within the service middleware.
func newUsersServiceTestFixture() *usersServiceTestFixture {
	mockTokenService := new(services.MockTokenService)
	mockRepository := new(persistence.MockUsersRepository)
	mockSecurityService := new(services.MockSecurityService)

	updatedEmail := "stub updated email"
	updatedUsername := "stub updated username"
	updatedPassword := "stub updated password"
	updatedBio := "stub updated bio"
	updatedImage := "stub updated image"
	stubUpdateUserRequest.Email = &updatedEmail
	stubUpdateUserRequest.Username = &updatedUsername
	stubUpdateUserRequest.Email = &updatedPassword
	stubUpdateUserRequest.Email = &updatedBio
	stubUpdateUserRequest.Email = &updatedImage

	return &usersServiceTestFixture{
		mockTokenService:    mockTokenService,
		mockRepository:      mockRepository,
		mockSecurityService: mockSecurityService,
		service:             core.NewUsersService(nil, mockRepository, mockTokenService, mockSecurityService),
		ctx:                 context.Background(),
	}
}
