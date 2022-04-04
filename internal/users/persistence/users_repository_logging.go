package persistence

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"time"
)

type usersRepositoryLoggingMiddleware struct {
	logger log.Logger
	next   UsersRepository
}

func NewUsersRepositoryLoggingMiddleware(logger log.Logger) UsersRepositoryMiddleware {
	return func(next UsersRepository) UsersRepository {
		return &usersRepositoryLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *usersRepositoryLoggingMiddleware) GetUser(ctx context.Context, userId int) (user *UserEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetUser",
			"query_time", time.Since(begin),
			"user_id", userId,
			"error", err,
		)
	}(time.Now())

	return mw.next.GetUser(ctx, userId)
}

func (mw *usersRepositoryLoggingMiddleware) CreateUser(ctx context.Context, username string, email string, password string) (user *UserEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateUser",
			"email", email,
			"username", username,
			"query_time", time.Since(begin),
			"created", fmt.Sprint(user != nil),
			"error", err,
		)
	}(time.Now())

	return mw.next.CreateUser(ctx, username, email, password)
}

func (mw *usersRepositoryLoggingMiddleware) UpdateUser(ctx context.Context, userId int, username, email, password, bio, image string) (user *UserEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "UpdateUser",
			"email", email,
			"username", username,
			"bio", bio,
			"image", image,
			"query_time", time.Since(begin),
			"user_id", userId,
			"error", err,
		)
	}(time.Now())

	return mw.next.UpdateUser(ctx, userId, username, email, password, bio, image)
}

func (mw *usersRepositoryLoggingMiddleware) FindUserByUsername(ctx context.Context, username string) (user *UserEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "FindUserByUsername",
			"username", username,
			"query_time", time.Since(begin),
			"found", fmt.Sprint(user != nil),
			"error", err,
		)
	}(time.Now())

	return mw.next.FindUserByUsername(ctx, username)
}

func (mw *usersRepositoryLoggingMiddleware) FindUserByEmail(ctx context.Context, email string) (user *UserEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "FindUserByEmail",
			"email", email,
			"query_time", time.Since(begin),
			"found", fmt.Sprint(user != nil),
			"error", err,
		)
	}(time.Now())

	return mw.next.FindUserByEmail(ctx, email)
}

func (mw *usersRepositoryLoggingMiddleware) FindUserByUsernameOrEmail(ctx context.Context, username string, email string) (user *UserEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "FindUserByUsernameOrEmail",
			"email", email,
			"username", username,
			"query_time", time.Since(begin),
			"found", fmt.Sprint(user != nil),
			"error", err,
		)
	}(time.Now())

	return mw.next.FindUserByUsernameOrEmail(ctx, username, email)
}

func (mw *usersRepositoryLoggingMiddleware) GetUserProfileFollowByFollowee(ctx context.Context, followerUserId, followeeUserId int) (userProfileFollow *UserProfileFollowEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetUserProfileFollowByFollowee",
			"followerUserId", followerUserId,
			"followeeUserId", followeeUserId,
			"query_time", time.Since(begin),
			"found", fmt.Sprint(userProfileFollow != nil),
			"error", err,
		)
	}(time.Now())

	return mw.next.GetUserProfileFollowByFollowee(ctx, followerUserId, followeeUserId)
}

func (mw *usersRepositoryLoggingMiddleware) CreateUserFollow(ctx context.Context, followerUserId, followeeUserId int) (userFollow *UserProfileFollowEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateUserFollow",
			"followerUserId", followerUserId,
			"followeeUserId", followeeUserId,
			"query_time", time.Since(begin),
			"follow_id", userFollow.Id,
			"error", err,
		)
	}(time.Now())

	return mw.next.CreateUserFollow(ctx, followerUserId, followeeUserId)
}

func (mw *usersRepositoryLoggingMiddleware) RemoveUserFollow(ctx context.Context, followerUserId, followeeUserId int) (err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "RemoveUserFollow",
			"followerUserId", followerUserId,
			"followeeUserId", followeeUserId,
			"query_time", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return mw.next.RemoveUserFollow(ctx, followerUserId, followeeUserId)
}
