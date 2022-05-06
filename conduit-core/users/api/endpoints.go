package api

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users/core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users/domain"
    "github.com/joeymckenzie/realworld-go-kit/pkg/api"
    "github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

func makeRegisterUserEndpoint(service core.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(domain.RegisterUserApiRequest)
        response, err := service.RegisterUser(ctx, &domain.RegisterUserServiceRequest{
            Email:    apiRequest.User.Email,
            Username: apiRequest.User.Username,
            Password: apiRequest.User.Password,
        })

        if err != nil {
            return nil, err
        }

        apiResponse := domain.UserResponse{
            User: *response,
        }

        return apiResponse, nil
    }
}

func makeLoginUserEndpoint(service core.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(domain.LoginUserApiRequest)
        response, err := service.LoginUser(ctx, &domain.LoginUserServiceRequest{
            Email:    apiRequest.User.Email,
            Password: apiRequest.User.Password,
        })

        if err != nil {
            return nil, err
        }

        apiResponse := domain.UserResponse{
            User: *response,
        }

        return apiResponse, nil
    }
}

func makeGetUserEndpoint(service core.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            response, err := service.GetUser(ctx, tokenMeta.UserId)

            if err != nil {
                return nil, err
            }

            apiResponse := domain.UserResponse{
                User: *response,
            }

            return apiResponse, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeUpdateUserEndpoint(service core.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            apiRequest := request.(domain.UpdateUserApiRequest)
            serviceRequest := &domain.UpdateUserServiceRequest{
                UserId:   tokenMeta.UserId,
                Email:    apiRequest.User.Email,
                Username: apiRequest.User.Username,
                Password: apiRequest.User.Password,
                Image:    apiRequest.User.Image,
                Bio:      apiRequest.User.Bio,
            }

            response, err := service.UpdateUser(ctx, serviceRequest)

            if err != nil {
                return nil, err
            }

            apiResponse := domain.UserResponse{
                User: *response,
            }

            return apiResponse, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeGetUserProfileEndpoint(service core.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(domain.GetUserProfileApiRequest)
        response, err := service.GetUserProfile(ctx, apiRequest.ProfileUsername, apiRequest.CurrentUserId)

        if err != nil {
            return nil, err
        }

        return domain.ProfileResponse{
            Profile: *response,
        }, nil
    }
}

func makeAddUserFollowEndpoint(service core.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            response, err := service.AddUserFollow(ctx, tokenMeta.UserId, request.(string))

            if err != nil {
                return nil, err
            }

            return domain.ProfileResponse{
                Profile: *response,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeRemoveUserFollowEndpoint(service core.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            response, err := service.RemoveUserFollow(ctx, tokenMeta.UserId, request.(string))

            if err != nil {
                return nil, err
            }

            return domain.ProfileResponse{
                Profile: *response,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}
