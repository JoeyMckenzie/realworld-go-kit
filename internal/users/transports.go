package users

import (
	"context"
	"encoding/json"
	"github.com/go-kit/log"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func MakeUserRoutes(logger log.Logger, router *chi.Mux, service UsersService) *chi.Mux {
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

	router.Post("/users", registerUserHandler.ServeHTTP)
	router.Post("/user/login", loginUserHandler.ServeHTTP)

	return router
}

func decodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AuthenticationRequest[RegisterUserRequest]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, shared.ErrInvalidRequestBody
	}

	return request, nil
}

func decodeLoginUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AuthenticationRequest[LoginUserRequest]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, shared.ErrInvalidRequestBody
	}

	return request, nil
}
