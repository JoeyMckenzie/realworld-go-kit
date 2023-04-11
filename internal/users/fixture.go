package users

import (
	"context"
	"github.com/go-kit/log"
)

var fixture *usersServiceTestFixture

type usersServiceTestFixture struct {
	ctx                 context.Context
	service             UsersService
	mockSecurityService *mockSecurityService
	mockTokenService    *mockTokenService
	mockRepository      *mockUsersRepository
}

func newUsersServiceTestFixture() *usersServiceTestFixture {
	ctx := context.Background()
	nopLogger := log.NewNopLogger()
	mockTokenService := new(mockTokenService)
	mockSecurityService := new(mockSecurityService)
	mockRepository := new(mockUsersRepository)
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
