package specs

import (
    "context"
    "entgo.io/ent/dialect"
    "github.com/joeymckenzie/realworld-go-kit/ent"
    "github.com/joeymckenzie/realworld-go-kit/ent/enttest"
    "github.com/joeymckenzie/realworld-go-kit/internal"
    "github.com/joeymckenzie/realworld-go-kit/internal/users/core"
    "github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
    "github.com/joeymckenzie/realworld-go-kit/pkg/services"
    _ "github.com/mattn/go-sqlite3"
    "testing"
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
    client              *ent.Client
    service             core.UsersService
    ctx                 context.Context
}

// newUsersServiceTestFixture sets up a common test fixture with in-place mocks for users service dependencies.
// Note that we don't need a validator dependency as validation is done within the service middleware.
func newUsersServiceTestFixture(t *testing.T) *usersServiceTestFixture {
    mockTokenService := new(services.MockTokenService)
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

    ctx := context.Background()
    testClient := enttest.Open(t, dialect.SQLite, "file:realworld_go_kit?mode=memory&cache=shared&_fk=1")
    internal.SeedData(ctx, testClient)

    return &usersServiceTestFixture{
        mockTokenService:    mockTokenService,
        mockSecurityService: mockSecurityService,
        client:              testClient,
        service:             core.NewUsersService(nil, testClient, mockTokenService, mockSecurityService),
        ctx:                 ctx,
    }
}
