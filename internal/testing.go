package internal

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/joeymckenzie/realworld-go-kit/ent"
)

func SetupTestFixture() (context.Context, *ent.Client) {
	ctx := context.Background()

	client, _ := ent.Open(dialect.SQLite, "file:realworld_go_kit?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	client.Schema.Create(ctx)

	SeedData(ctx, client)

	return ctx, client
}
