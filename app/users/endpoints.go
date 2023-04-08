package users

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeRegisterUserEndpoint(service UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		registrationRequest := request.(AuthenticationRequest)

		if registrationRequest.User.Email == nil {
			return nil, errors.New("bad request")
		}

		return nil, nil
	}
}
