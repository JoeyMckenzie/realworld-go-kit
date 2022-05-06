package specs

import (
    "context"
    "entgo.io/ent/dialect"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/persistence"
    _ "github.com/mattn/go-sqlite3"
    "os"
    "testing"
)

var fixture *commentsServiceTestFixture

type commentsServiceTestFixture struct {
    ctx     context.Context
    service core.CommentsService
}

func newCommentsServiceTestFixture(ctx context.Context, client *ent.Client) *commentsServiceTestFixture {
    return &commentsServiceTestFixture{
        ctx:     ctx,
        service: core.NewCommentsService(nil, client),
    }
}

func TestMain(m *testing.M) {
    ctx := context.Background()

    // Setup our in-memory database for ent
    client, _ := ent.Open(dialect.SQLite, "file:realworld_go_kit?mode=memory&cache=shared&_fk=1")

    defer client.Close()
    client.Schema.Create(ctx)

    // Seed test data and create the test fixture
    persistence.SeedData(ctx, client)
    fixture = newCommentsServiceTestFixture(ctx, client)

    // Finally, run our tests
    os.Exit(m.Run())
}
