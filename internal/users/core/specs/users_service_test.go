package specs

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/joeymckenzie/realworld-go-kit/internal"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var fixture *usersServiceTestFixture

type usersServiceTestFixture struct {
	ctx                 context.Context
	service             core.UsersService
	mockSecurityService *services.MockSecurityService
	mockTokenService    *services.MockTokenService
}

func NewUsersServiceTestFixture(ctx context.Context, client *ent.Client) *usersServiceTestFixture {
	mockTokenService := new(services.MockTokenService)
	mockSecurityService := new(services.MockSecurityService)

	return &usersServiceTestFixture{
		ctx:                 ctx,
		mockTokenService:    mockTokenService,
		mockSecurityService: mockSecurityService,
		service:             core.NewUsersService(nil, client, mockTokenService, mockSecurityService),
	}
}

func (usersServiceTestFixture) resetMocks() {
	fixture.mockTokenService.ResetMocks()
	fixture.mockSecurityService.ResetMocks()
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	client, _ := ent.Open(dialect.SQLite, "file:realworld_go_kit?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	client.Schema.Create(ctx)

	internal.SeedData(ctx, client)
	fixture = NewUsersServiceTestFixture(ctx, client)

	os.Exit(m.Run())
}
