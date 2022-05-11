package utilities

import (
    "context"
    "encoding/json"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-kit/kit/transport"
    httpTransport "github.com/go-kit/kit/transport/http"
    "github.com/go-kit/log"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "net/http"
    "time"
)

func NewChiRouter() *chi.Mux {
    router := chi.NewRouter()
    router.Use(CorsMiddleware)
    router.Use(JsonContentTypeMiddleware)
    router.Use(middleware.RequestID)
    router.Use(middleware.Timeout(30 * time.Second))

    return router
}

func DecodeDefaultRequest(_ context.Context, _ *http.Request) (interface{}, error) {
    return nil, nil
}

func EncodeSuccessfulResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    if _, ok := response.(error); ok {
        w.WriteHeader(http.StatusBadRequest)
    }

    return json.NewEncoder(w).Encode(response)
}

func EncodeSuccessfulResponseWithNoContent(_ context.Context, w http.ResponseWriter, response interface{}) error {
    if _, ok := response.(error); ok {
        w.WriteHeader(http.StatusBadRequest)
    }

    w.WriteHeader(http.StatusNoContent)

    return nil
}

func HandlerOptions(logger log.Logger) []httpTransport.ServerOption {
    return []httpTransport.ServerOption{
        httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
        httpTransport.ServerErrorEncoder(EncodeError),
    }
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
    if err == nil {
        // Note: we have bigger problems if this happens...
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(shared.NewGenericError())
        return
    }

    // On unauthorized, don't provide any context for security and hand back 401
    if err == utilities.ErrUnauthorized {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    if apiError, ok := err.(*shared.ConduitError); ok {
        w.WriteHeader(apiError.Code)
        json.NewEncoder(w).Encode(apiError)
    } else {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(shared.NewGenericError())
    }
}
