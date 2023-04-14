package middleware

import (
    "context"
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "github.com/joeymckenzie/realworld-go-kit/internal/users/core"
)

type usersServiceValidationMiddleware struct {
    validation *validator.Validate
    next       core.UsersService
}

func NewUsersServiceValidationMiddleware(validation *validator.Validate) core.UsersServiceMiddleware {
    return func(next core.UsersService) core.UsersService {
        return &usersServiceValidationMiddleware{
            validation: validation,
            next:       next,
        }
    }
}

func (mw *usersServiceValidationMiddleware) Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (*domain.User, error) {
    if err := mw.validation.StructCtx(ctx, request); err != nil {
        return &domain.User{}, shared.MakeValidationError(err)
    }

    return mw.next.Register(ctx, request)
}

func (mw *usersServiceValidationMiddleware) Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (*domain.User, error) {
    if err := mw.validation.StructCtx(ctx, request); err != nil {
        return &domain.User{}, shared.MakeValidationError(err)
    }

    return mw.next.Login(ctx, request)
}

func (mw *usersServiceValidationMiddleware) Update(ctx context.Context, request domain.AuthenticationRequest[domain.UpdateUserRequest], id uuid.UUID) (*domain.User, error) {
    return mw.next.Update(ctx, request, id)
}

func (mw *usersServiceValidationMiddleware) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
    return mw.next.Get(ctx, id)
}
