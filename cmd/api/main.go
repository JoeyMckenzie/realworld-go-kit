package main

import (
	"fmt"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal"
	"github.com/joeymckenzie/realworld-go-kit/pkg/config"
	"github.com/joeymckenzie/realworld-go-kit/pkg/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
	"os"
)

func main() {
	logger := utilities.InitializeLogger()
	defer level.Info(logger).Log("main", "application shutting down...")

	environment, port := config.InitializeConfiguration(logger)
	entClient := persistence.InitializeEnt(logger, environment)
	usersService, articlesService := internal.InitializeServices(logger, entClient)
	router := internal.InitializeRouter(logger, usersService, articlesService)

	serverPort := fmt.Sprintf(":%d", port)
	level.Info(logger).Log("server_start", fmt.Sprintf("listening on port %s", serverPort))

	if err := http.ListenAndServe(serverPort, router); err != nil {
		level.Error(logger).Log("main", "error during server startup", "error", err)
		os.Exit(1)
	}
}
