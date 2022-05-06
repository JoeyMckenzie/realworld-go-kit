package config

import (
	"flag"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"os"
)

func InitializeConfiguration(logger log.Logger) (string, int, bool) {
	// Load in environment variables
	if err := godotenv.Load(); err != nil {
		level.Error(logger).Log(
			"environment", "error while bootstrapping environment",
			"error", err,
		)
		os.Exit(1)
	}

	// Load in configuration
	environment := flag.String("env", "development", "Environment to run the application under")
	port := flag.Int("port", 8080, "Environment to run the application under")
	applySeed := flag.Bool("seed", false, "Seed the database on startup")
	flag.Parse()

	if environment == nil || *environment == "" {
		level.Error(logger).Log("environment", "no environment provided at startup")
		os.Exit(1)
	}

	if port == nil {
		level.Error(logger).Log("port", "no port provided at startup")
		os.Exit(1)
	}

	if applySeed == nil {
		seed := false
		applySeed = &seed
	}

	return *environment, *port, *applySeed
}
