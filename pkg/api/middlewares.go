package api

import (
	"context"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
	"github.com/rs/cors"
	"net/http"
)

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	corsConfig := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodGet},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         3600,
	}

	return cors.New(corsConfig).Handler(next)
}

type TokenMeta struct {
	UserId   int
	Username string
	Email    string
}

func AuthorizedRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := services.
			NewTokenService().
			GetRequiredUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

		if err != nil || userId < 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		requestContext := context.WithValue(r.Context(), TokenMeta{}, TokenMeta{UserId: userId})
		r = r.WithContext(requestContext)
		next.ServeHTTP(w, r)
	})
}
