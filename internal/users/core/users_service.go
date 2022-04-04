package core

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
)

type (
	UsersService interface {
		RegisterUser(ctx context.Context, user *domain.RegisterUserServiceRequest) (*domain.UserDto, error)
		LoginUser(ctx context.Context, userRequest *domain.LoginUserServiceRequest) (*domain.UserDto, error)
		GetUser(ctx context.Context, userId int) (*domain.UserDto, error)
		UpdateUser(ctx context.Context, request *domain.UpdateUserServiceRequest) (*domain.UserDto, error)
		GetUserProfile(ctx context.Context, username string, currentUserId int) (*domain.ProfileDto, error)
		AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*domain.ProfileDto, error)
		RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*domain.ProfileDto, error)
	}

	usersService struct {
		validator       *validator.Validate
		repository      persistence.UsersRepository
		tokenService    services.TokenService
		securityService services.SecurityService
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewUsersService(validator *validator.Validate, repository persistence.UsersRepository, tokenService services.TokenService, securityService services.SecurityService) UsersService {
	return &usersService{
		validator:       validator,
		repository:      repository,
		tokenService:    tokenService,
		securityService: securityService,
	}
}

func (us *usersService) GetUser(ctx context.Context, userId int) (*domain.UserDto, error) {
	// Retrieve the mapped user, returning any service utilities that occur
	existingUser, err := us.repository.GetUser(ctx, userId)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	// Generate a new JWT for the user
	token, err := us.tokenService.GenerateUserToken(existingUser.Id, existingUser.Email)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	return domain.NewUserDto(existingUser.Email, existingUser.Username, existingUser.Bio, existingUser.Image, token), nil
}

func (us *usersService) RegisterUser(ctx context.Context, request *domain.RegisterUserServiceRequest) (*domain.UserDto, error) {
	// Verify the username and password are available
	existingUser, err := us.repository.FindUserByUsernameOrEmail(ctx, request.Username, request.Email)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if existingUser != nil {
		return nil, api.NewApiErrorWithContext(http.StatusConflict, "user", utilities.ErrUsernameOrEmailTaken)
	}

	// Hash the user password with bcrypt
	hashedPassword, err := us.securityService.HashPassword(request.Password)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	// Insert the user, propagate any errors as 500s
	createdUser, err := us.repository.CreateUser(ctx, request.Username, request.Email, hashedPassword)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	// Generate a JWT for the user
	token, err := us.tokenService.GenerateUserToken(createdUser.Id, createdUser.Email)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	return domain.NewDefaultUserDto(createdUser.Email, createdUser.Username, token), nil
}

func (us *usersService) LoginUser(ctx context.Context, request *domain.LoginUserServiceRequest) (*domain.UserDto, error) {
	// Verify the username and password are available
	existingUser, err := us.repository.FindUserByEmail(ctx, request.Email)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, api.NewApiErrorWithContext(http.StatusConflict, "user", utilities.ErrUserNotFound)
	}

	// Compare password hashes for identity
	if passwordIsValid := us.securityService.PasswordIsValid(existingUser.Password, request.Password); !passwordIsValid {
		return nil, api.NewApiErrorWithContext(http.StatusUnauthorized, "user", utilities.ErrInvalidLoginAttempt)
	}

	// Generate a JWT for the user
	token, err := us.tokenService.GenerateUserToken(existingUser.Id, existingUser.Email)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	return domain.NewUserDto(existingUser.Email, existingUser.Username, existingUser.Bio, existingUser.Image, token), nil
}

func (us *usersService) UpdateUser(ctx context.Context, request *domain.UpdateUserServiceRequest) (*domain.UserDto, error) {
	// Verify the existing user, return unauthorized for obfuscation
	existingUser, err := us.repository.GetUser(ctx, request.UserId)
	userNotFound := existingUser == nil || err == sql.ErrNoRows
	if err != nil || userNotFound || existingUser.Id != request.UserId {
		return nil, utilities.ErrUnauthorized
	}

	updatedUsername := updateIfRequired(existingUser.Username, request.Username)
	updatedEmail := updateIfRequired(existingUser.Email, request.Email)
	updatedBio := updateIfRequired(existingUser.Bio, request.Bio)
	updatedImage := updateIfRequired(existingUser.Image, request.Image)
	updatedPassword := existingUser.Password

	if request.Password != nil {
		updatedHashedPassword, err := us.securityService.HashPassword(*request.Password)

		if err != nil {
			return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
		}

		updatedPassword = updatedHashedPassword
	}

	// Retrieve the mapped user, returning any service utilities that occur
	user, err := us.repository.UpdateUser(ctx, existingUser.Id, updatedUsername, updatedEmail, updatedPassword, updatedBio, updatedImage)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusBadRequest, "user", err)
	}

	// Generate a new JWT for the user
	token, err := us.tokenService.GenerateUserToken(existingUser.Id, updatedEmail)
	if err != nil {
		return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	return domain.NewUserDto(user.Email, user.Username, user.Bio, user.Image, token), nil
}

func (us *usersService) GetUserProfile(ctx context.Context, username string, currentUserId int) (*domain.ProfileDto, error) {
	existingUser, err := us.repository.FindUserByUsername(ctx, username)

	if isValidDatabaseErr(err) {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, api.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	isFollowing := false

	if currentUserId > 0 {
		existingUserFollow, err := us.repository.GetUserProfileFollowByFollowee(ctx, currentUserId, existingUser.Id)

		if isValidDatabaseErr(err) {
			return nil, err
		}

		if existingUserFollow != nil {
			isFollowing = true
		}
	}

	return &domain.ProfileDto{
		Username:  existingUser.Username,
		Bio:       existingUser.Bio,
		Image:     existingUser.Image,
		Following: isFollowing,
	}, nil
}

func (us *usersService) AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*domain.ProfileDto, error) {
	userToFollow, err := us.repository.FindUserByUsername(ctx, followeeUsername)

	if isValidDatabaseErr(err) {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, api.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	if _, err = us.repository.CreateUserFollow(ctx, followerUserId, userToFollow.Id); err != nil {
		return nil, api.NewGenericError()
	}

	return &domain.ProfileDto{
		Username:  userToFollow.Username,
		Bio:       userToFollow.Bio,
		Image:     userToFollow.Bio,
		Following: true,
	}, nil
}

func (us *usersService) RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*domain.ProfileDto, error) {
	userToFollow, err := us.repository.FindUserByUsername(ctx, followeeUsername)

	if isValidDatabaseErr(err) {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, api.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	if err = us.repository.RemoveUserFollow(ctx, followerUserId, userToFollow.Id); err != nil {
		return nil, api.NewGenericError()
	}

	return &domain.ProfileDto{
		Username:  userToFollow.Username,
		Bio:       userToFollow.Bio,
		Image:     userToFollow.Bio,
		Following: false,
	}, nil
}
