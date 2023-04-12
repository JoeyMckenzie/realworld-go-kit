package core

import (
	"context"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/infrastructure"
)

var fixture *usersServiceTestFixture

type usersServiceTestFixture struct {
	ctx                 context.Context
	service             UsersService
	mockSecurityService *infrastructure.MockSecurityService
	mockTokenService    *shared.MockTokenService
	mockRepository      *infrastructure.MockUsersRepository
}

func newUsersServiceTestFixture() *usersServiceTestFixture {
	ctx := context.Background()
	nopLogger := log.NewNopLogger()
	mockTokenService := new(shared.MockTokenService)
	mockSecurityService := new(infrastructure.MockSecurityService)
	mockRepository := new(infrastructure.MockUsersRepository)
	service := NewService(nopLogger, mockRepository, mockTokenService, mockSecurityService)

	return &usersServiceTestFixture{
		ctx:                 ctx,
		service:             service,
		mockTokenService:    mockTokenService,
		mockSecurityService: mockSecurityService,
		mockRepository:      mockRepository,
	}
}

func (f *usersServiceTestFixture) resetMocks() {
	f.mockTokenService.ResetMocks()
	f.mockSecurityService.ResetMocks()
	f.mockRepository.ResetMocks()
}
