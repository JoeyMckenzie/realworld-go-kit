package api

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
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

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		// Note: we have bigger problems if this happens...
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewGenericError())
		return
	}

	// On unauthorized, don't provide any context for security and hand back 401
	if err == utilities.ErrUnauthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	statusCode := http.StatusInternalServerError

	if apiError, ok := err.(*ApiErrors); ok {
		statusCode = apiError.Code
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(struct {
			Errors ApiError `json:"errors"`
		}{
			Errors: apiError.Errors,
		})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewGenericError())
	}
}
