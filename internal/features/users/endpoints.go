package users

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/utilities"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func makeRegisterUserEndpoint(service UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        registrationRequest := request.(domain.AuthenticationRequest[domain.RegisterUserRequest])
        createdUser, err := service.Register(ctx, registrationRequest)

        if err != nil {
            return nil, shared.ErrorWithContext("error while registering in user", err)
        }

        return &domain.AuthenticationResponse{
            User: createdUser,
        }, nil
    }
}

func makeLoginUserEndpoint(service UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        loginRequest := request.(domain.AuthenticationRequest[domain.LoginUserRequest])
        verifiedUser, err := service.Login(ctx, loginRequest)

        if err != nil {
            return nil, shared.ErrorWithContext("error while logging in user", err)
        }

        return &domain.AuthenticationResponse{
            User: verifiedUser,
        }, nil
    }
}

func makeUpdateUserEndpoint(service UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
        updateRequest := request.(domain.AuthenticationRequest[domain.UpdateUserRequest])
        updatedUser, err := service.Update(ctx, updateRequest, uuidClaim.UserId)

        if err != nil {
            return nil, shared.ErrorWithContext("error while updating in user", err)
        }

        return &domain.AuthenticationResponse{
            User: updatedUser,
        }, nil
    }
}

func makeGetUserEndpoint(service UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
        existingUser, err := service.Get(ctx, uuidClaim.UserId)

        if err != nil {
            return nil, shared.ErrorWithContext("error while retrieving in user", err)
        }

        return &domain.AuthenticationResponse{
            User: existingUser,
        }, nil
    }
}
