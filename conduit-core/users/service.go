package users

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
	sharedDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/shared"
	usersDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/users"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/follow"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/user"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"net/http"
	"time"
)

type (
	UsersService interface {
		RegisterUser(ctx context.Context, user *usersDomain.RegisterUserServiceRequest) (*sharedDomain.UserDto, error)
		LoginUser(ctx context.Context, userRequest *usersDomain.LoginUserServiceRequest) (*sharedDomain.UserDto, error)
		GetUser(ctx context.Context, userId int) (*sharedDomain.UserDto, error)
		UpdateUser(ctx context.Context, request *usersDomain.UpdateUserServiceRequest) (*sharedDomain.UserDto, error)
		GetUserProfile(ctx context.Context, username string, currentUserId int) (*sharedDomain.ProfileDto, error)
		AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*sharedDomain.ProfileDto, error)
		RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*sharedDomain.ProfileDto, error)
	}

	usersService struct {
		validator       *validator.Validate
		client          *ent.Client
		tokenService    services.TokenService
		securityService services.SecurityService
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewUsersService(validator *validator.Validate, client *ent.Client, tokenService services.TokenService, securityService services.SecurityService) UsersService {
	return &usersService{
		validator:       validator,
		client:          client,
		tokenService:    tokenService,
		securityService: securityService,
	}
}

func (us *usersService) GetUser(ctx context.Context, userId int) (*sharedDomain.UserDto, error) {
	// Retrieve the mapped user, returning any service utilities that occur
	existingUser, err := us.client.User.Get(ctx, userId)

	if ent.IsNotFound(err) {
		return nil, shared.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	// Generate a new JWT for the user
	token, err := us.tokenService.GenerateUserToken(userId, existingUser.Email)
	if err != nil {
		return nil, shared.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	return usersDomain.NewUserDto(existingUser.Email, existingUser.Username, existingUser.Bio, existingUser.Image, token), nil
}

func (us *usersService) RegisterUser(ctx context.Context, request *usersDomain.RegisterUserServiceRequest) (*sharedDomain.UserDto, error) {
	// Verify the username and password are available
	existingUsers, err := us.client.User.
		Query().
		Where(
			user.Or(
				user.Username(request.Username),
				user.Email(request.Email),
			),
		).
		All(ctx)

	if err != nil {
		// return nil, shared.NewInternalServerErrorWithContext("user", err)
		return nil, err
	} else if len(existingUsers) > 0 {
		return nil, shared.NewApiErrorWithContext(http.StatusConflict, "user", utilities.ErrUsernameOrEmailTaken)
	}

	// Hash the user password with bcrypt
	hashedPassword, err := us.securityService.HashPassword(request.Password)

	if err != nil {
		return nil, shared.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	// Insert the user, propagate any errors as 500s
	createdUser, err := us.client.User.
		Create().
		SetUsername(request.Username).
		SetEmail(request.Email).
		SetPassword(hashedPassword).
		SetNillableImage(nil).
		SetNillableBio(nil).
		Save(ctx)

	if err != nil {
		return nil, shared.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	// Generate a JWT for the user
	token, err := us.tokenService.GenerateUserToken(createdUser.ID, createdUser.Email)

	if err != nil {
		return nil, shared.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	return usersDomain.NewDefaultUserDto(createdUser.Email, createdUser.Username, token), nil
}

func (us *usersService) LoginUser(ctx context.Context, request *usersDomain.LoginUserServiceRequest) (*sharedDomain.UserDto, error) {
	// Verify the exists
	existingUser, err := us.client.User.
		Query().
		Where(user.Email(request.Email)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, shared.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	// Compare password hashes for identity
	if passwordIsValid := us.securityService.IsValidPassword(existingUser.Password, request.Password); !passwordIsValid {
		return nil, shared.NewApiErrorWithContext(http.StatusUnauthorized, "user", utilities.ErrInvalidLoginAttempt)
	}

	// Generate a JWT for the user
	token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return nil, shared.NewInternalServerErrorWithContext("user", err)
	}

	return usersDomain.NewUserDto(existingUser.Email, existingUser.Username, existingUser.Bio, existingUser.Image, token), nil
}

func (us *usersService) UpdateUser(ctx context.Context, request *usersDomain.UpdateUserServiceRequest) (*sharedDomain.UserDto, error) {
	// Verify the existing user, return unauthorized for obfuscation
	existingUser, err := us.client.User.Get(ctx, request.UserId)

	if ent.IsNotFound(err) {
		return nil, utilities.ErrUnauthorized
	}

	updatedUsername := utilities.UpdateIfRequired(existingUser.Username, request.Username)
	updatedEmail := utilities.UpdateIfRequired(existingUser.Email, request.Email)
	updatedBio := utilities.UpdateIfRequired(existingUser.Bio, request.Bio)
	updatedImage := utilities.UpdateIfRequired(existingUser.Image, request.Image)
	updatedPassword := existingUser.Password

	if request.Password != nil {
		updatedHashedPassword, err := us.securityService.HashPassword(*request.Password)

		if err != nil {
			return nil, shared.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
		}

		updatedPassword = updatedHashedPassword
	}

	// Retrieve the mapped user, returning any service utilities that occur
	updatedUser, err := us.client.User.
		UpdateOne(existingUser).
		SetUsername(updatedUsername).
		SetEmail(updatedEmail).
		SetPassword(updatedPassword).
		SetBio(updatedBio).
		SetImage(updatedImage).
		SetUpdateTime(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, shared.NewApiErrorWithContext(http.StatusBadRequest, "user", err)
	}

	// Generate a new JWT for the user
	token, err := us.tokenService.GenerateUserToken(existingUser.ID, updatedEmail)
	if err != nil {
		return nil, shared.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
	}

	return usersDomain.NewUserDto(updatedUser.Email, updatedUser.Username, updatedUser.Bio, updatedUser.Image, token), nil
}

func (us *usersService) GetUserProfile(ctx context.Context, username string, currentUserId int) (*sharedDomain.ProfileDto, error) {
	existingUser, err := us.client.User.
		Query().
		Where(user.Username(username)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, shared.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	isFollowing := false

	if currentUserId > 0 {
		_, err := us.client.Follow.
			Query().
			Where(
				follow.FollowerID(currentUserId),
				follow.FolloweeID(existingUser.ID),
			).
			First(ctx)

		if !ent.IsNotFound(err) {
			isFollowing = true
		}
	}

	return &sharedDomain.ProfileDto{
		Username:  existingUser.Username,
		Bio:       existingUser.Bio,
		Image:     existingUser.Image,
		Following: isFollowing,
	}, nil
}

func (us *usersService) AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*sharedDomain.ProfileDto, error) {
	// Verify both the user follower and followees exist
	userToFollow, err := us.client.User.
		Query().
		Where(user.Username(followeeUsername)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, shared.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	// Verify the user is not trying to follow themselves (common... you're better than that)
	if userToFollow.ID == followerUserId {
		return nil, shared.NewApiErrorWithContext(http.StatusBadRequest, "follows", utilities.ErrCannotFollowSelf)
	}

	if _, err = us.client.Follow.
		Create().
		SetFollowerID(followerUserId).
		SetFolloweeID(userToFollow.ID).
		Save(ctx); err != nil {
		return nil, shared.NewInternalServerErrorWithContext("follow", err)
	}

	return &sharedDomain.ProfileDto{
		Username:  userToFollow.Username,
		Bio:       userToFollow.Bio,
		Image:     userToFollow.Bio,
		Following: true,
	}, nil
}

func (us *usersService) RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*sharedDomain.ProfileDto, error) {
	userToUnfollow, err := us.client.User.
		Query().
		Where(user.Username(followeeUsername)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, shared.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
	}

	// Only propagate errors if something other than not found occurs,
	// as we don't _really_ care if someone tries to unfollow someone
	// that they're not following in the first place
	if _, err = us.client.Follow.
		Delete().
		Where(
			follow.And(
				follow.FollowerID(followerUserId),
				follow.FolloweeID(userToUnfollow.ID),
			),
		).
		Exec(ctx); err != nil && !ent.IsNotFound(err) {
		return nil, shared.NewInternalServerErrorWithContext("unfollow", err)
	}

	return &sharedDomain.ProfileDto{
		Username:  userToUnfollow.Username,
		Bio:       userToUnfollow.Bio,
		Image:     userToUnfollow.Bio,
		Following: false,
	}, nil
}
