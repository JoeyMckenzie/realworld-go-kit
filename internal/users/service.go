package users

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"net/http"
)

type (
	UsersService interface {
		Register(ctx context.Context, request AuthenticationRequest[RegisterUserRequest]) (*User, error)
	}

	userService struct {
		logger     log.Logger
		repository UsersRepository
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewService(logger log.Logger, repository UsersRepository) UsersService {
	return &userService{
		logger:     logger,
		repository: repository,
	}
}

func (us *userService) Register(ctx context.Context, request AuthenticationRequest[RegisterUserRequest]) (*User, error) {
	level.Info(us.logger).Log("registration", "attempting to register new user", "username", request.User.Username, "email", request.User.Email)

	// First, retrieve the user to verify the username and email is available
	existingUsers, err := us.repository.SearchUsers(ctx, *request.User.Username, *request.User.Email)

	if len(existingUsers) != 0 {
		// Technically, there could be at least two entries here - only log out the first one
		existingUser := existingUsers[0]
		level.Error(us.logger).Log("registration", "username or email is taken", "username", existingUser.Username, "email", existingUser.Email)
		return nil, shared.MakeApiError(http.StatusConflict, shared.ErrUsernameOrEmailTaken)
	} else if err != nil {
		level.Error(us.logger).Log("registration", "error while attempting to query for existing users", "err", err)
		return nil, shared.ErrInternalServerError
	}

	// Register the user while returning the entire mapped user back from the insert
	createdUser, err := us.repository.CreateUser(ctx, *request.User.Username, *request.User.Email, *request.User.Password)

	if err != nil {
		level.Error(us.logger).Log("registration", "error while attempting create user", "err", err)
		return nil, err
	}

	level.Info(us.logger).Log("registration", "user successfully created", "username", createdUser.Username, "email", createdUser.Email, "user_id", createdUser.ID.String())

	return createdUser.ToUser("token"), nil
}
