package core

import (
    "context"
    "github.com/go-kit/log"
    "github.com/joeymckenzie/realworld-go-kit/internal/users/infrastructure"
    "github.com/joeymckenzie/realworld-go-kit/internal/utilities"
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
    mockSecurityService *utilities.MockSecurityService
    mockTokenService    *utilities.MockTokenService
    mockRepository      *infrastructure.MockUsersRepository
}

func newUsersServiceTestFixture() *usersServiceTestFixture {
    ctx := context.Background()
    nopLogger := log.NewNopLogger()
    mockTokenService := new(utilities.MockTokenService)
    mockSecurityService := new(utilities.MockSecurityService)
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
