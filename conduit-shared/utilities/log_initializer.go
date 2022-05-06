package utilities

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"os"
)

func InitializeLogger() log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "caller", log.DefaultCaller)
		level.Info(logger).Log("main", "bootstrapping application...")
	}

	return logger
}
