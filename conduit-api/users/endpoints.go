package users

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    apiUtilities "github.com/joeymckenzie/realworld-go-kit/conduit-api/utilities"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users"
    usersDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/users"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
)

func makeRegisterUserEndpoint(service users.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(usersDomain.RegisterUserApiRequest)
        response, err := service.RegisterUser(ctx, &usersDomain.RegisterUserServiceRequest{
            Email:    apiRequest.User.Email,
            Username: apiRequest.User.Username,
            Password: apiRequest.User.Password,
        })

        if err != nil {
            return nil, err
        }

        apiResponse := usersDomain.UserResponse{
            User: *response,
        }

        return apiResponse, nil
    }
}

func makeLoginUserEndpoint(service users.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(usersDomain.LoginUserApiRequest)
        response, err := service.LoginUser(ctx, &usersDomain.LoginUserServiceRequest{
            Email:    apiRequest.User.Email,
            Password: apiRequest.User.Password,
        })

        if err != nil {
            return nil, err
        }

        apiResponse := usersDomain.UserResponse{
            User: *response,
        }

        return apiResponse, nil
    }
}

func makeGetUserEndpoint(service users.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(apiUtilities.TokenMeta{}).(apiUtilities.TokenMeta); ok && tokenMeta.UserId > 0 {
            response, err := service.GetUser(ctx, tokenMeta.UserId)

            if err != nil {
                return nil, err
            }

            apiResponse := usersDomain.UserResponse{
                User: *response,
            }

            return apiResponse, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeUpdateUserEndpoint(service users.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(apiUtilities.TokenMeta{}).(apiUtilities.TokenMeta); ok && tokenMeta.UserId > 0 {
            apiRequest := request.(usersDomain.UpdateUserApiRequest)
            serviceRequest := &usersDomain.UpdateUserServiceRequest{
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

            apiResponse := usersDomain.UserResponse{
                User: *response,
            }

            return apiResponse, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeGetUserProfileEndpoint(service users.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(usersDomain.GetUserProfileApiRequest)
        response, err := service.GetUserProfile(ctx, apiRequest.ProfileUsername, apiRequest.CurrentUserId)

        if err != nil {
            return nil, err
        }

        return usersDomain.ProfileResponse{
            Profile: *response,
        }, nil
    }
}

func makeAddUserFollowEndpoint(service users.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(apiUtilities.TokenMeta{}).(apiUtilities.TokenMeta); ok && tokenMeta.UserId > 0 {
            response, err := service.AddUserFollow(ctx, tokenMeta.UserId, request.(string))

            if err != nil {
                return nil, err
            }

            return usersDomain.ProfileResponse{
                Profile: *response,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeRemoveUserFollowEndpoint(service users.UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(apiUtilities.TokenMeta{}).(apiUtilities.TokenMeta); ok && tokenMeta.UserId > 0 {
            response, err := service.RemoveUserFollow(ctx, tokenMeta.UserId, request.(string))

            if err != nil {
                return nil, err
            }

            return usersDomain.ProfileResponse{
                Profile: *response,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}
