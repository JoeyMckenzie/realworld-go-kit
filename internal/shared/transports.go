package shared

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/exp/slog"
)

func DecodeNilPayload(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeSuccessfulOkResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if _, ok := response.(error); ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	if response != nil {
		return json.NewEncoder(w).Encode(response)
	}

	return nil
}

func HandlerOptions(logger *slog.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorHandler(newLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(EncodeError),
	}
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		// Note: we have bigger problems if this happens...
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MakeGenericApiError())
		return
	}

	// On unauthorized, don't provide any context for security and hand back 401
	if err == ErrUnauthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Attempt to parse the error as an API error, else return a generic error back to the user
	// If we type check the error as a validation error, add a 422 error code and serialize the errors
	if apiError, ok := err.(*ApiError[string]); ok {
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
	} else if validationError, ok := err.(*ApiError[[]string]); ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(validationError)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MakeGenericApiError())
	}
}
