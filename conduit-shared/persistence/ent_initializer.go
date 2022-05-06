package persistence

import (
    "context"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
    "entgo.io/ent/dialect/sql/schema"
    "entgo.io/ent/entc/integration/ent"
    "entgo.io/ent/entc/integration/ent/migrate"
    "fmt"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    _ "github.com/lib/pq"
    "os"
)

func InitializeEnt(logger log.Logger, environment string, applySeed bool) (*ent.Client, *sql.Driver, error) {
    // Build out our connection to the database
    var connectionString string
    {
        // Our API running within a docker context will need to communicate to the postgres container within the swarm
        if environment == "docker" {
            connectionString = os.Getenv("CONNECTION_STRING_DOCKER")
        } else {
            connectionString = os.Getenv("CONNECTION_STRING")
        }
    }

    // Generate the ent client
    driver, err := sql.Open(dialect.Postgres, connectionString)

    if err != nil {
        return nil, nil, err
    }

    driverWithDebugContext := dialect.DebugWithContext(driver, func(ctx context.Context, i ...interface{}) {
        level.Debug(logger).Log("query", fmt.Sprintf("%v", i))
    })

    entClient := ent.NewClient(ent.Driver(driverWithDebugContext))

    // Run the auto migration tool
    ctx := context.Background()
    err = entClient.Schema.Create(
        context.Background(),
        schema.WithAtlas(true),
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true))

    if applySeed {
        SeedData(ctx, entClient)
    }

    if err != nil {
        level.Error(logger).Log("main", "failed running auto migrations", "error", err)
        os.Exit(1)
    }

    return entClient, driver, nil
}
