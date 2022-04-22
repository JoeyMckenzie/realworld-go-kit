package specs

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/joeymckenzie/realworld-go-kit/internal"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var fixture *articlesServiceTestFixture

type articlesServiceTestFixture struct {
	ctx     context.Context
	service core.ArticlesService
}

func newArticlesServiceTestFixture(ctx context.Context, client *ent.Client) *articlesServiceTestFixture {
	return &articlesServiceTestFixture{
		ctx:     ctx,
		service: core.NewArticlesServices(nil, client),
	}
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Setup our in-memory database for ent
	client, _ := ent.Open(dialect.SQLite, "file:realworld_go_kit?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	client.Schema.Create(ctx)

	// Seed test data and create the test fixture
	internal.SeedData(ctx, client)
	fixture = newArticlesServiceTestFixture(ctx, client)

	// Finally, run our tests
	os.Exit(m.Run())
}
