package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeRegisterUserEndpoint(service UsersService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        registrationRequest := request.(AuthenticationRequest)
        createdUser, err := service.Register(ctx, registrationRequest)

        if err != nil {
            return nil, err
        }

        return createdUser, nil
    }
}
