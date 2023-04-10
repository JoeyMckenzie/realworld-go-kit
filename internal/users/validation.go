package users

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type usersServiceValidationMiddleware struct {
	validation *validator.Validate
	next       UsersService
}

func NewUsersServiceValidationMiddleware(validation *validator.Validate) UsersServiceMiddleware {
	return func(next UsersService) UsersService {
		return &usersServiceValidationMiddleware{
			validation: validation,
			next:       next,
		}
	}
}

func (mw *usersServiceValidationMiddleware) Register(ctx context.Context, request AuthenticationRequest) (user *User, err error) {
	if validationErrors := mw.validation.StructCtx(ctx, request); err != nil {
		return &User{}, shared.MakeValidationError(validationErrors)
	}

	return mw.next.Register(ctx, request)
}
