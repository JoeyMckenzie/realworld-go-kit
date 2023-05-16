package shared

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gofrs/uuid"
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

		requestContext := context.WithValue(r.Context(), TokenContextKey{}, TokenContextKey{UserId: userId})
		r = r.WithContext(requestContext)
		next.ServeHTTP(w, r)
	})
}

func UsernameRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		if username == "" {
			w.WriteHeader(http.StatusNotFound)
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
			// Okay if we don't find a user on this optionally authenticated route, default to the nil value
			userId = uuid.Nil
		}

		requestContext := context.WithValue(r.Context(), TokenContextKey{}, TokenContextKey{UserId: userId})
		r = r.WithContext(requestContext)
		next.ServeHTTP(w, r)
	})
}
