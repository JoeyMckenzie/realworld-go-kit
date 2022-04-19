package core

import (
    "context"
    "github.com/go-playground/validator/v10"
    "github.com/joeymckenzie/realworld-go-kit/ent"
    "github.com/joeymckenzie/realworld-go-kit/ent/follow"
    "github.com/joeymckenzie/realworld-go-kit/ent/user"
    "github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
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

func (us *usersService) GetUser(ctx context.Context, userId int) (*domain.UserDto, error) {
    // Retrieve the mapped user, returning any service utilities that occur
    existingUser, err := us.client.User.Get(ctx, userId)
    if err != nil {
        return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
    } else if ent.IsNotFound(err) {
        return nil, api.NewApiErrorWithContext(http.StatusNotFound, "user", err)
    }

    // Generate a new JWT for the user
    token, err := us.tokenService.GenerateUserToken(userId, existingUser.Email)
    if err != nil {
        return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
    }

    return domain.NewUserDto(existingUser.Email, existingUser.Username, existingUser.Bio, existingUser.Image, token), nil
}

func (us *usersService) RegisterUser(ctx context.Context, request *domain.RegisterUserServiceRequest) (*domain.UserDto, error) {
    // Verify the username and password are available
    existingUsers, err := us.client.User.
        Query().
        Where(
            user.Email(request.Email),
            user.Username(request.Username)).
        All(ctx)

    if err != nil && !ent.IsNotFound(err) {
        return nil, err
    } else if len(existingUsers) > 0 {
        return nil, api.NewApiErrorWithContext(http.StatusConflict, "user", utilities.ErrUsernameOrEmailTaken)
    }

    // Hash the user password with bcrypt
    hashedPassword, err := us.securityService.HashPassword(request.Password)
    if err != nil {
        return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
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
        return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
    }

    // Generate a JWT for the user
    token, err := us.tokenService.GenerateUserToken(createdUser.ID, createdUser.Email)
    if err != nil {
        return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
    }

    return domain.NewDefaultUserDto(createdUser.Email, createdUser.Username, token), nil
}

func (us *usersService) LoginUser(ctx context.Context, request *domain.LoginUserServiceRequest) (*domain.UserDto, error) {
    // Verify the exists
    existingUser, err := us.client.User.
        Query().
        Where(user.Email(request.Email)).
        First(ctx)

    if err != nil && !ent.IsNotFound(err) {
        return nil, err
    } else if ent.IsNotFound(err) {
        return nil, api.NewApiErrorWithContext(http.StatusConflict, "user", utilities.ErrUserNotFound)
    }

    // Compare password hashes for identity
    if passwordIsValid := us.securityService.PasswordIsValid(existingUser.Password, request.Password); !passwordIsValid {
        return nil, api.NewApiErrorWithContext(http.StatusUnauthorized, "user", utilities.ErrInvalidLoginAttempt)
    }

    // Generate a JWT for the user
    token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)
    if err != nil {
        return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
    }

    return domain.NewUserDto(existingUser.Email, existingUser.Username, existingUser.Bio, existingUser.Image, token), nil
}

func (us *usersService) UpdateUser(ctx context.Context, request *domain.UpdateUserServiceRequest) (*domain.UserDto, error) {
    // Verify the existing user, return unauthorized for obfuscation
    existingUser, err := us.client.User.Get(ctx, request.UserId)
    if err != nil || existingUser.ID != request.UserId {
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
    updatedUser, err := us.client.User.
        UpdateOneID(existingUser.ID).
        SetUsername(updatedUsername).
        SetEmail(updatedEmail).
        SetPassword(updatedPassword).
        SetBio(updatedBio).
        SetImage(updatedImage).
        Save(ctx)

    if err != nil {
        return nil, api.NewApiErrorWithContext(http.StatusBadRequest, "user", err)
    }

    // Generate a new JWT for the user
    token, err := us.tokenService.GenerateUserToken(existingUser.ID, updatedEmail)
    if err != nil {
        return nil, api.NewApiErrorWithContext(http.StatusInternalServerError, "user", err)
    }

    return domain.NewUserDto(updatedUser.Email, updatedUser.Username, updatedUser.Bio, updatedUser.Image, token), nil
}

func (us *usersService) GetUserProfile(ctx context.Context, username string, currentUserId int) (*domain.ProfileDto, error) {
    existingUser, err := us.client.User.
        Query().
        Where(user.Username(username)).
        First(ctx)

    if isValidDatabaseErr(err) {
        return nil, api.NewInternalServerErrorWithContext("user", err)
    } else if ent.IsNotFound(err) {
        return nil, api.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
    }

    isFollowing := false

    if currentUserId > 0 {
        _, err := us.client.User.
            QueryFollowers(existingUser).
            Where(follow.FolloweeID(currentUserId)).
            First(ctx)

        if isValidDatabaseErr(err) {
            return nil, api.NewInternalServerErrorWithContext("user", err)
        }

        if !ent.IsNotFound(err) {
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
    userToFollow, err := us.client.User.
        Query().
        Where(user.Username(followeeUsername)).
        First(ctx)

    if isValidDatabaseErr(err) {
        return nil, api.NewInternalServerErrorWithContext("user", err)
    } else if ent.IsNotFound(err) {
        return nil, api.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
    }

    // Verify the user is not trying to follow themselves (common... you're better than that)
    if userToFollow.ID == followerUserId {
        return nil, api.NewApiErrorWithContext(http.StatusBadRequest, "follows", utilities.ErrCannotFollowSelf)
    }

    _, err = us.client.Follow.
        Create().
        SetFollowerID(followerUserId).
        SetFolloweeID(userToFollow.ID).
        Save(ctx)

    if _, err = us.client.Follow.
        Create().
        SetFollowerID(followerUserId).
        SetFolloweeID(userToFollow.ID).
        Save(ctx); err != nil {
        return nil, api.NewInternalServerErrorWithContext("follow", err)
    }

    return &domain.ProfileDto{
        Username:  userToFollow.Username,
        Bio:       userToFollow.Bio,
        Image:     userToFollow.Bio,
        Following: true,
    }, nil
}

func (us *usersService) RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*domain.ProfileDto, error) {
    userToUnfollow, err := us.client.User.
        Query().
        Where(user.Username(followeeUsername)).
        First(ctx)

    if isValidDatabaseErr(err) {
        return nil, err
    } else if ent.IsNotFound(err) {
        return nil, api.NewApiErrorWithContext(http.StatusNotFound, "user", utilities.ErrUserNotFound)
    }

    if _, err = us.client.Follow.
        Delete().
        Where(
            follow.FollowerID(followerUserId),
            follow.FolloweeID(userToUnfollow.ID)).
        Exec(ctx); err != nil {
        return nil, api.NewInternalServerErrorWithContext("follow", err)
    }

    return &domain.ProfileDto{
        Username:  userToUnfollow.Username,
        Bio:       userToUnfollow.Bio,
        Image:     userToUnfollow.Bio,
        Following: false,
    }, nil
}
