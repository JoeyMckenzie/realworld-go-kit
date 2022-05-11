package users

import (
    "context"
    "encoding/json"
    "github.com/go-chi/chi/v5"
    httpTransport "github.com/go-kit/kit/transport/http"
    "github.com/go-kit/log"
    apiUtilities "github.com/joeymckenzie/realworld-go-kit/conduit-api/utilities"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users"
    usersDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/users"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "net/http"
)

func MakeUsersTransport(router *chi.Mux, logger log.Logger, service users.UsersService) *chi.Mux {
    registerUserHandler := httpTransport.NewServer(
        makeRegisterUserEndpoint(service),
        decodeRegisterUserRequest,
        apiUtilities.EncodeSuccessfulResponse,
        apiUtilities.HandlerOptions(logger)...,
    )

    loginUserHandler := httpTransport.NewServer(
        makeLoginUserEndpoint(service),
        decodeLoginUserRequest,
        apiUtilities.EncodeSuccessfulResponse,
        apiUtilities.HandlerOptions(logger)...,
    )

    getCurrentUserHandler := httpTransport.NewServer(
        makeGetUserEndpoint(service),
        apiUtilities.DecodeDefaultRequest,
        apiUtilities.EncodeSuccessfulResponse,
        apiUtilities.HandlerOptions(logger)...,
    )

    getUserProfileHandler := httpTransport.NewServer(
        makeGetUserProfileEndpoint(service),
        decodeGetUserProfileRequest,
        apiUtilities.EncodeSuccessfulResponse,
        apiUtilities.HandlerOptions(logger)...,
    )

    updateUserHandler := httpTransport.NewServer(
        makeUpdateUserEndpoint(service),
        decodeUpdateUserRequest,
        apiUtilities.EncodeSuccessfulResponse,
        apiUtilities.HandlerOptions(logger)...,
    )

    addUserFollowHandler := httpTransport.NewServer(
        makeAddUserFollowEndpoint(service),
        decodeUserFollowRequest,
        apiUtilities.EncodeSuccessfulResponse,
        apiUtilities.HandlerOptions(logger)...,
    )

    removeUserFollowHandler := httpTransport.NewServer(
        makeRemoveUserFollowEndpoint(service),
        decodeUserFollowRequest,
        apiUtilities.EncodeSuccessfulResponse,
        apiUtilities.HandlerOptions(logger)...,
    )

    router.Route("/profiles", func(r chi.Router) {
        r.Get("/{username}", getUserProfileHandler.ServeHTTP)

        // Authorized profile requests for following/unfollowing usersDomain
        r.Group(func(r chi.Router) {
            r.Use(apiUtilities.AuthorizedRequestMiddleware)
            r.Post("/{username}/follow", addUserFollowHandler.ServeHTTP)
            r.Delete("/{username}/follow", removeUserFollowHandler.ServeHTTP)
        })
    })

    // Login/register handlers
    router.Route("/usersDomain", func(r chi.Router) {
        r.Post("/", registerUserHandler.ServeHTTP)
        r.Post("/login", loginUserHandler.ServeHTTP)
    })

    // Authenticated usersDomain requests flows for updating and retrieving user information
    router.Route("/user", func(r chi.Router) {
        r.Use(apiUtilities.AuthorizedRequestMiddleware)
        r.Put("/", updateUserHandler.ServeHTTP)
        r.Get("/", getCurrentUserHandler.ServeHTTP)
    })

    return router
}

func decodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request usersDomain.RegisterUserApiRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, utilities.ErrInvalidRequestBody
    }

    return request, nil
}

func decodeLoginUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request usersDomain.LoginUserApiRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, utilities.ErrInvalidRequestBody
    }

    return request, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request usersDomain.UpdateUserApiRequest

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

    return usersDomain.GetUserProfileApiRequest{
        CurrentUserId:   userId,
        ProfileUsername: username,
    }, nil
}

func decodeUserFollowRequest(_ context.Context, r *http.Request) (interface{}, error) {
    return chi.URLParam(r, "username"), nil
}
