package users

import (
	"context"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/data"
	utilities2 "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/utilities"
)

const (
	stubEmail    = "email@email.com"
	stubUsername = "username"
	stubPassword = "password"
	stubImage    = "image"
	stubBio      = "bio"
)

var fixture *usersServiceTestFixture

type usersServiceTestFixture struct {
	ctx                 context.Context
	service             UsersService
	mockSecurityService *utilities2.MockSecurityService
	mockTokenService    *utilities2.MockTokenService
	mockRepository      *data.MockUsersRepository
}

func newUsersServiceTestFixture() *usersServiceTestFixture {
	ctx := context.Background()
	nopLogger := log.NewNopLogger()
	mockTokenService := new(utilities2.MockTokenService)
	mockSecurityService := new(utilities2.MockSecurityService)
	mockRepository := new(data.MockUsersRepository)
	service := NewUsersService(nopLogger, mockRepository, mockTokenService, mockSecurityService)

	return &usersServiceTestFixture{
		ctx:                 ctx,
		service:             service,
		mockTokenService:    mockTokenService,
		mockSecurityService: mockSecurityService,
		mockRepository:      mockRepository,
	}
}
