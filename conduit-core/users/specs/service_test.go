package specs

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/users"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/persistence"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var fixture *usersServiceTestFixture

type usersServiceTestFixture struct {
    ctx                 context.Context
    service             users.UsersService
    mockSecurityService *services.MockSecurityService
    mockTokenService    *services.MockTokenService
}

func newUsersServiceTestFixture(ctx context.Context, client *ent.Client) *usersServiceTestFixture {
    mockTokenService := new(services.MockTokenService)
    mockSecurityService := new(services.MockSecurityService)

    return &usersServiceTestFixture{
        ctx:                 ctx,
        mockTokenService:    mockTokenService,
        mockSecurityService: mockSecurityService,
        service:             users.NewUsersService(nil, client, mockTokenService, mockSecurityService),
    }
}

func (usersServiceTestFixture) resetMocks() {
    fixture.mockTokenService.ResetMocks()
    fixture.mockSecurityService.ResetMocks()
}

func TestMain(m *testing.M) {
    ctx := context.Background()

    // Setup our in-memory database for ent
    client, _ := ent.Open(dialect.SQLite, "file:realworld_go_kit?mode=memory&cache=shared&_fk=1")
    defer client.Close()
    client.Schema.Create(ctx)

    // Seed test data and create the test fixture
    persistence.SeedData(ctx, client)
    fixture = newUsersServiceTestFixture(ctx, client)

    // Finally, run our tests
    os.Exit(m.Run())
}
