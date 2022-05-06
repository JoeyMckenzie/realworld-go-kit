package main

import (
    "entgo.io/ent/dialect/sql"
    "fmt"
    "github.com/go-kit/log/level"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/config"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/persistence"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "net/http"
    "os"
)

func main() {
    logger := utilities.InitializeLogger()

    defer level.Info(logger).Log("main", "application shutting down...")

    environment, port, applySeed := config.InitializeConfiguration(logger)
    entClient, driver, err := persistence.InitializeEnt(logger, environment, applySeed)

    defer func(driver *sql.Driver) {
        if err := driver.Close(); err != nil {
            level.Error(logger).Log("main", "failed closing postgres connection", "error", err)
            os.Exit(1)
        }
    }(driver)

    if err != nil {
        level.Error(logger).Log("database_driver", err)
        os.Exit(1)
    }

    serviceRegister := conduit_core.InitializeServices(logger, entClient)
    router := conduit_core.InitializeRouter(logger, serviceRegister)
    serverPort := fmt.Sprintf(":%d", port)

    level.Info(logger).Log("server_start", fmt.Sprintf("listening on port %s", serverPort))

    if err := http.ListenAndServe(serverPort, router); err != nil {
        level.Error(logger).Log("main", "error during server startup", "error", err)
        os.Exit(1)
    }
}
