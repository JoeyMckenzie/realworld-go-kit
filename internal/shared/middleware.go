package shared

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joeymckenzie/realworld-go-kit/internal/utilities"
	"net/http"
)

func JsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func CorsPolicy(next http.Handler) http.Handler {
	corsConfig := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         3600,
	}

	return cors.New(corsConfig).Handler(next)
}

func AuthorizationRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, ok := GetUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		requestContext := context.WithValue(r.Context(), utilities.TokenContextKey{}, utilities.TokenContextKey{UserId: userId})
		r = r.WithContext(requestContext)
		next.ServeHTTP(w, r)
	})
}

func UsernameRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		if username == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		requestContext := context.WithValue(r.Context(), UsernameContextKey{}, UsernameContextKey{Username: username})
		r = r.WithContext(requestContext)
		next.ServeHTTP(w, r)
	})
}

func AuthorizationOptional(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, ok := GetUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

		if !ok {
			return
		}

		requestContext := context.WithValue(r.Context(), utilities.TokenContextKey{}, utilities.TokenContextKey{UserId: userId})
		r = r.WithContext(requestContext)
		next.ServeHTTP(w, r)
	})
}
