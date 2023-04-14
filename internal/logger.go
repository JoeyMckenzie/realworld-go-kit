package internal

import (
    "os"

    "github.com/go-kit/log"
)

// NewLogger spins up a kit logger for printing to stdout, though we could swap this out for something like axiom eventually.
func NewLogger() log.Logger {
    var logger log.Logger
    {
        logger = log.NewLogfmtLogger(os.Stderr)
        logger = log.NewSyncLogger(logger)
        logger = log.With(logger, "caller", log.DefaultCaller)
    }

    return logger
}
