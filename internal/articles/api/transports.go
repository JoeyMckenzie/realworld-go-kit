package api

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
)

func MakeArticlesHandler(logger log.Logger, validator *validator.Validate) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/articles", func(r chi.Router) {
	})

	return router
}

func decodeArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.UpsertArticleApiRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utilities.ErrInvalidRequestBody
	}

	return request, nil
}
