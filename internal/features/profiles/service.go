package profiles

import (
    "context"
    "database/sql"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

type (
    ProfilesService interface {
        GetProfile(ctx context.Context, username string, followeeId uuid.UUID) (*domain.Profile, error)
        Follow(ctx context.Context, username string, followeeId uuid.UUID) (*domain.Profile, error)
        Unfollow(ctx context.Context, username string, followeeId uuid.UUID) (*domain.Profile, error)
    }

    profileService struct {
        logger     log.Logger
        repository repositories.UsersRepository
    }

    ProfileServiceMiddleware func(service ProfilesService) ProfilesService
)

func NewProfileService(logger log.Logger, repository repositories.UsersRepository) ProfilesService {
    return &profileService{
        logger:     logger,
        repository: repository,
    }
}

func (us *profileService) GetProfile(ctx context.Context, username string, followeeId uuid.UUID) (*domain.Profile, error) {
    const loggingSpan string = "get_profile"

    level.Info(us.logger).Log(loggingSpan, "attempting to retrieve profile status, verifying existing user", "username", username, "followee_id", followeeId)
    existingUser, err := us.repository.GetUserByUsername(ctx, username)

    if err != nil && err != sql.ErrNoRows {
        level.Error(us.logger).Log(loggingSpan, "error while attempting to retrieve user profile", "username", username, "followee_id", followeeId, "err", err)
        return &domain.Profile{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        level.Error(us.logger).Log(loggingSpan, "user profile was not found", "username", username, "followee_id", followeeId)
        return &domain.Profile{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

    isFollowing := false

    if followeeId != uuid.Nil {
        followerId := existingUser.ID
        level.Info(us.logger).Log(loggingSpan, "checking for existing user follow", "username", username, "follower_id", followerId, "followee_id", followeeId)
        id, err := us.repository.GetExistingFollow(ctx, followerId, followeeId)

        if err != nil && err != sql.ErrNoRows {
            level.Error(us.logger).Log(loggingSpan, "error while attempting checking for existing user follow", "username", username, "follower_id", followerId, "followee_id", followeeId, "err", err)
            return &domain.Profile{}, shared.MakeApiError(err)
        } else if id != uuid.Nil {
            level.Warn(us.logger).Log(loggingSpan, "found existing follow for user", "username", username, "follower_id", followerId, "followee_id", followeeId)
            isFollowing = true
        }
    }

    level.Info(us.logger).Log(loggingSpan, "user follower successfully added", "follower_id", existingUser.ID, "followee_id", followeeId)

    return existingUser.ToProfile(isFollowing), nil
}

func (us *profileService) Follow(ctx context.Context, username string, followeeId uuid.UUID) (*domain.Profile, error) {
    const loggingSpan string = "follower_user"

    level.Info(us.logger).Log(loggingSpan, "attempting to add user follower, verifying existing user to follow", "username", username, "followee_id", followeeId)
    existingUserToFollow, err := us.repository.GetUserByUsername(ctx, username)

    if err != nil && err != sql.ErrNoRows {
        level.Error(us.logger).Log(loggingSpan, "error while attempting to search for user to follow", "username", username, "followee_id", followeeId, "err", err)
        return &domain.Profile{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        level.Error(us.logger).Log(loggingSpan, "user to follow was not found", "username", username, "followee_id", followeeId)
        return &domain.Profile{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

    followerId := existingUserToFollow.ID
    level.Info(us.logger).Log(loggingSpan, "checking for existing user follow", "username", username, "follower_id", followerId, "followee_id", followeeId)
    id, err := us.repository.GetExistingFollow(ctx, followerId, followeeId)

    if err != nil && err != sql.ErrNoRows {
        level.Error(us.logger).Log(loggingSpan, "error while attempting checking for existing user follow", "username", username, "follower_id", followerId, "followee_id", followeeId, "err", err)
        return &domain.Profile{}, shared.MakeApiError(err)
    } else if id != uuid.Nil {
        level.Warn(us.logger).Log(loggingSpan, "user follow already exists, skipping", "username", username, "follower_id", followerId, "followee_id", followeeId)
        return existingUserToFollow.ToProfile(true), nil
    }

    if err := us.repository.AddFollow(ctx, existingUserToFollow.ID, followeeId); err != nil {
        return &domain.Profile{}, err
    }

    level.Info(us.logger).Log(loggingSpan, "user follower successfully added", "follower_id", existingUserToFollow.ID, "followee_id", followeeId)

    return existingUserToFollow.ToProfile(true), nil
}

func (us *profileService) Unfollow(ctx context.Context, username string, followeeId uuid.UUID) (*domain.Profile, error) {
    const loggingSpan string = "unfollower_user"

    level.Info(us.logger).Log(loggingSpan, "attempting to delete user follower, verifying existing user to follow", "username", username, "followee_id", followeeId)
    existingUserToUnfollow, err := us.repository.GetUserByUsername(ctx, username)

    if err != nil && err != sql.ErrNoRows {
        level.Error(us.logger).Log(loggingSpan, "error while attempting to search for user to unfollow", "username", username, "followee_id", followeeId, "err", err)
        return &domain.Profile{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        level.Error(us.logger).Log(loggingSpan, "user to unfollow was not found", "username", username, "followee_id", followeeId)
        return &domain.Profile{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

    // No need to check for existing follows, we'll only delete if we find a match
    if err := us.repository.DeleteFollow(ctx, existingUserToUnfollow.ID, followeeId); err != nil {
        level.Error(us.logger).Log(loggingSpan, "error while attempting to delete user follow", "username", username, "followee_id", followeeId, "err", err)
        return &domain.Profile{}, shared.MakeApiError(err)
    }

    level.Info(us.logger).Log(loggingSpan, "user follower successfully removed", "follower_id", existingUserToUnfollow.ID, "followee_id", followeeId)

    return existingUserToUnfollow.ToProfile(false), nil
}
