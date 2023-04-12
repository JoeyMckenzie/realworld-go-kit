package core

import (
    "context"
    "github.com/go-playground/validator/v10"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "github.com/joeymckenzie/realworld-go-kit/internal/users"
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

func (mw *usersServiceValidationMiddleware) Register(ctx context.Context, request users.AuthenticationRequest[users.RegisterUserRequest]) (*users.User, error) {
    if err := mw.validation.StructCtx(ctx, request); err != nil {
        return &users.User{}, shared.MakeValidationError(err)
    }

    return mw.next.Register(ctx, request)
}

func (mw *usersServiceValidationMiddleware) Login(ctx context.Context, request users.AuthenticationRequest[users.LoginUserRequest]) (*users.User, error) {
    if err := mw.validation.StructCtx(ctx, request); err != nil {
        return &users.User{}, shared.MakeValidationError(err)
    }

    return mw.next.Login(ctx, request)
}

func (mw *usersServiceValidationMiddleware) Update(ctx context.Context, request users.AuthenticationRequest[users.UpdateUserRequest]) (*users.User, error) {
    return mw.next.Update(ctx, request)
}
