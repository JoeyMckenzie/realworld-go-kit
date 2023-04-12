package api

import (
	"context"
	"encoding/json"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func MakeUserRoutes(logger log.Logger, router *chi.Mux, service core.UsersService) *chi.Mux {
	registerUserHandler := httptransport.NewServer(
		makeRegisterUserEndpoint(service),
		decodeRegisterUserRequest,
		shared.EncodeSuccessfulResponse,
		shared.HandlerOptions(logger)...,
	)

	loginUserHandler := httptransport.NewServer(
		makeLoginUserEndpoint(service),
		decodeLoginUserRequest,
		shared.EncodeSuccessfulResponse,
		shared.HandlerOptions(logger)...,
	)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", registerUserHandler.ServeHTTP)
		r.Post("/login", loginUserHandler.ServeHTTP)
	})

	return router
}

func decodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request users.AuthenticationRequest[users.RegisterUserRequest]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, shared.ErrInvalidRequestBody
	}

	return request, nil
}

func decodeLoginUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request users.AuthenticationRequest[users.LoginUserRequest]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, shared.ErrInvalidRequestBody
	}

	return request, nil
}
