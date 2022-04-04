package api

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

func makeRegisterUserEndpoint(service core.UsersService, validator *validator.Validate) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if apiRequest, ok := request.(domain.RegisterUserApiRequest); ok {
			if err := validator.Struct(apiRequest); err != nil {
				return nil, api.NewValidationError(err)
			}

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

		return nil, utilities.ErrInternalServerError
	}
}

func makeLoginUserEndpoint(service core.UsersService, validator *validator.Validate) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if apiRequest, ok := request.(domain.LoginUserApiRequest); ok {
			if err := validator.Struct(apiRequest); err != nil {
				return nil, api.NewValidationError(err)
			}

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

		return nil, utilities.ErrInternalServerError
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

func makeUpdateUserEndpoint(service core.UsersService, validator *validator.Validate) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if apiRequest, ok := request.(domain.UpdateUserApiRequest); ok {
			if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
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
			} else {
				return nil, utilities.ErrUnauthorized
			}
		}

		return nil, utilities.ErrInternalServerError
	}
}

func makeGetUserProfileEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if apiRequest, ok := request.(domain.GetUserProfileApiRequest); ok {
			response, err := service.GetUserProfile(ctx, apiRequest.ProfileUsername, apiRequest.CurrentUserId)

			if err != nil {
				return nil, err
			}

			return domain.ProfileResponse{
				Profile: *response,
			}, nil
		}

		return nil, utilities.ErrInternalServerError
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

		return nil, utilities.ErrInternalServerError
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

		return nil, utilities.ErrInternalServerError
	}
}
