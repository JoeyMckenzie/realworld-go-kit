package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeRegisterUserEndpoint(service UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		registrationRequest := request.(AuthenticationRequest[RegisterUserRequest])
		createdUser, err := service.Register(ctx, registrationRequest)

		if err != nil {
			return nil, err
		}

		return &AuthenticationResponse{
			User: createdUser,
		}, nil
	}
}
