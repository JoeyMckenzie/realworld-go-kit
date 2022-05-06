package api

import (
    "context"
    "encoding/json"
    httpTransport "github.com/go-kit/kit/transport/http"
    "github.com/go-kit/log"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users/core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users/domain"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/api"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "net/http"
)

func MakeUsersTransport(router *chi.Mux, logger log.Logger, service core.UsersService) *chi.Mux {
    registerUserHandler := httpTransport.NewServer(
        makeRegisterUserEndpoint(service),
        decodeRegisterUserRequest,
        api.EncodeSuccessfulResponse,
        api.HandlerOptions(logger)...,
    )

    loginUserHandler := httpTransport.NewServer(
        makeLoginUserEndpoint(service),
        decodeLoginUserRequest,
        api.EncodeSuccessfulResponse,
        api.HandlerOptions(logger)...,
    )

    getCurrentUserHandler := httpTransport.NewServer(
        makeGetUserEndpoint(service),
        api.DecodeDefaultRequest,
        api.EncodeSuccessfulResponse,
        api.HandlerOptions(logger)...,
    )

    getUserProfileHandler := httpTransport.NewServer(
        makeGetUserProfileEndpoint(service),
        decodeGetUserProfileRequest,
        api.EncodeSuccessfulResponse,
        api.HandlerOptions(logger)...,
    )

    updateUserHandler := httpTransport.NewServer(
        makeUpdateUserEndpoint(service),
        decodeUpdateUserRequest,
        api.EncodeSuccessfulResponse,
        api.HandlerOptions(logger)...,
    )

    addUserFollowHandler := httpTransport.NewServer(
        makeAddUserFollowEndpoint(service),
        decodeUserFollowRequest,
        api.EncodeSuccessfulResponse,
        api.HandlerOptions(logger)...,
    )

    removeUserFollowHandler := httpTransport.NewServer(
        makeRemoveUserFollowEndpoint(service),
        decodeUserFollowRequest,
        api.EncodeSuccessfulResponse,
        api.HandlerOptions(logger)...,
    )

    router.Route("/profiles", func(r chi.Router) {
        r.Get("/{username}", getUserProfileHandler.ServeHTTP)

        // Authorized profile requests for following/unfollowing users
        r.Group(func(r chi.Router) {
            r.Use(api.AuthorizedRequestMiddleware)
            r.Post("/{username}/follow", addUserFollowHandler.ServeHTTP)
            r.Delete("/{username}/follow", removeUserFollowHandler.ServeHTTP)
        })
    })

    // Login/register handlers
    router.Route("/users", func(r chi.Router) {
        r.Post("/", registerUserHandler.ServeHTTP)
        r.Post("/login", loginUserHandler.ServeHTTP)
    })

    // Authenticated users requests flows for updating and retrieving user information
    router.Route("/user", func(r chi.Router) {
        r.Use(api.AuthorizedRequestMiddleware)
        r.Put("/", updateUserHandler.ServeHTTP)
        r.Get("/", getCurrentUserHandler.ServeHTTP)
    })

    return router
}

func decodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request domain.RegisterUserApiRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, utilities.ErrInvalidRequestBody
    }

    return request, nil
}

func decodeLoginUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request domain.LoginUserApiRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, utilities.ErrInvalidRequestBody
    }

    return request, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request domain.UpdateUserApiRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, utilities.ErrInvalidRequestBody
    }

    return request, nil
}

func decodeGetUserProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
    username := chi.URLParam(r, "username")
    userId, err := services.
        NewTokenService().
        GetOptionalUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

    if err != nil {
        return nil, utilities.ErrInternalServerError
    }

    return domain.GetUserProfileApiRequest{
        CurrentUserId:   userId,
        ProfileUsername: username,
    }, nil
}

func decodeUserFollowRequest(_ context.Context, r *http.Request) (interface{}, error) {
    return chi.URLParam(r, "username"), nil
}
