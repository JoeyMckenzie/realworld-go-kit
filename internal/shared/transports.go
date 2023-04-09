package shared

import (
	"context"
	"encoding/json"
	"net/http"
)

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
