package api

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/utilities"
)

func makeRegisterUserEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		registrationRequest := request.(users.AuthenticationRequest[users.RegisterUserRequest])
		createdUser, err := service.Register(ctx, registrationRequest)

		if err != nil {
			return nil, err
		}

		return &users.AuthenticationResponse{
			User: createdUser,
		}, nil
	}
}

func makeLoginUserEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		loginRequest := request.(users.AuthenticationRequest[users.LoginUserRequest])
		verifiedUser, err := service.Login(ctx, loginRequest)

		if err != nil {
			return nil, err
		}

		return &users.AuthenticationResponse{
			User: verifiedUser,
		}, nil
	}
}

func makeUpdateUserEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
		updateRequest := request.(users.AuthenticationRequest[users.UpdateUserRequest])
		updatedUser, err := service.Update(ctx, updateRequest, uuidClaim.UserId)

		if err != nil {
			return nil, err
		}

		return &users.AuthenticationResponse{
			User: updatedUser,
		}, nil
	}
}

func makeGetUserEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
		existingUser, err := service.Get(ctx, uuidClaim.UserId)

		if err != nil {
			return nil, err
		}

		return &users.AuthenticationResponse{
			User: existingUser,
		}, nil
	}
}

func makeGetProfileEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
		usernameToFollow := ctx.Value(shared.UsernameContextKey{}).(shared.UsernameContextKey)
		profile, err := service.GetProfile(ctx, usernameToFollow.Username, uuidClaim.UserId)

		if err != nil {
			return nil, err
		}

		return &users.ProfileResponse{
			Profile: profile,
		}, nil
	}
}

func makeFollowUserEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
		usernameToFollow := ctx.Value(shared.UsernameContextKey{}).(shared.UsernameContextKey)
		profile, err := service.Follow(ctx, usernameToFollow.Username, uuidClaim.UserId)

		if err != nil {
			return nil, err
		}

		return &users.ProfileResponse{
			Profile: profile,
		}, nil
	}
}

func makeUnfollowUserEndpoint(service core.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
		usernameToUnfollow := ctx.Value(shared.UsernameContextKey{}).(shared.UsernameContextKey)
		profile, err := service.Unfollow(ctx, usernameToUnfollow.Username, uuidClaim.UserId)

		if err != nil {
			return nil, err
		}

		return &users.ProfileResponse{
			Profile: profile,
		}, nil
	}
}
