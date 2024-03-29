package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"golang.org/x/exp/slog"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func MakeUserRoutes(logger *slog.Logger, router *chi.Mux, service UsersService) *chi.Mux {
	registerUserHandler := httptransport.NewServer(
		makeRegisterUserEndpoint(service),
		decodeRegisterUserRequest,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	loginUserHandler := httptransport.NewServer(
		makeLoginUserEndpoint(service),
		decodeLoginUserRequest,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	updateUserHandler := httptransport.NewServer(
		makeUpdateUserEndpoint(service),
		decodeUpdateUserRequest,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	getUserHandler := httptransport.NewServer(
		makeGetUserEndpoint(service),
		shared.DecodeNilPayload,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", registerUserHandler.ServeHTTP)
		r.Post("/login", loginUserHandler.ServeHTTP)
	})

	router.Route("/user", func(r chi.Router) {
		r.Use(shared.AuthorizationRequired)
		r.Put("/", updateUserHandler.ServeHTTP)
		r.Get("/", getUserHandler.ServeHTTP)
	})

	return router
}

func decodeRegisterUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return decodeUserRequest[domain.RegisterUserRequest](ctx, r)
}

func decodeLoginUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return decodeUserRequest[domain.LoginUserRequest](ctx, r)
}

func decodeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return decodeUserRequest[domain.UpdateUserRequest](ctx, r)
}

func decodeUserRequest[T domain.RegisterUserRequest | domain.LoginUserRequest | domain.UpdateUserRequest](_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.AuthenticationRequest[T]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, shared.ErrorWithContext("error while attempting to decode the authentication request", shared.ErrInvalidRequestBody)
	}

	return request, nil
}
