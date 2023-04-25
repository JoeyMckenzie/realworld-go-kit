package shared

import (
    "context"
    "golang.org/x/exp/slog"
)

type logErrorHandler struct {
    logger *slog.Logger
}

func (h logErrorHandler) Handle(ctx context.Context, err error) {
    h.logger.ErrorCtx(ctx, "err", err)
}

func newLogErrorHandler(logger *slog.Logger) *logErrorHandler {
    return &logErrorHandler{
        logger: logger,
    }
}
