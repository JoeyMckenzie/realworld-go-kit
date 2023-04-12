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

    updateUserHandler := httptransport.NewServer(
        makeUpdateUserEndpoint(service),
        decodeUpdateUserRequest,
        shared.EncodeSuccessfulResponse,
        shared.HandlerOptions(logger)...,
    )

    router.Route("/users", func(r chi.Router) {
        r.Post("/", registerUserHandler.ServeHTTP)
        r.Post("/login", loginUserHandler.ServeHTTP)
    })

    router.Put("/user", updateUserHandler.ServeHTTP)

    return router
}

func decodeRegisterUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    return decodeUserRequest[users.RegisterUserRequest](ctx, r)
}

func decodeLoginUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    return decodeUserRequest[users.LoginUserRequest](ctx, r)
}

func decodeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    return decodeUserRequest[users.UpdateUserRequest](ctx, r)
}

func decodeUserRequest[T users.RegisterUserRequest | users.LoginUserRequest | users.UpdateUserRequest](_ context.Context, r *http.Request) (interface{}, error) {
    var request users.AuthenticationRequest[T]

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, shared.ErrInvalidRequestBody
    }

    return request, nil
}
