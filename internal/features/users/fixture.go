package users

import (
	"context"

	"github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
	"github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/utilities"
	"golang.org/x/exp/slog"
)

var fixture *usersServiceTestFixture

type usersServiceTestFixture struct {
	ctx                 context.Context
	service             UsersService
	mockSecurityService *utilities.MockSecurityService
	mockTokenService    *utilities.MockTokenService
	mockRepository      *repositories.MockUsersRepository
}

func newUsersServiceTestFixture() *usersServiceTestFixture {
	ctx := context.Background()
	nopLogger := slog.Default()
	mockTokenService := new(utilities.MockTokenService)
	mockSecurityService := new(utilities.MockSecurityService)
	mockRepository := new(repositories.MockUsersRepository)
	service := NewUsersService(nopLogger, mockRepository, mockTokenService, mockSecurityService)

	return &usersServiceTestFixture{
		ctx:                 ctx,
		service:             service,
		mockTokenService:    mockTokenService,
		mockSecurityService: mockSecurityService,
		mockRepository:      mockRepository,
	}
}
