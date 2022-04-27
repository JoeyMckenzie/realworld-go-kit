package middlewares

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/domain"
	"time"
)

type usersServiceMetricsMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	service        core.UsersService
}

func NewUsersServiceMetricsMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) core.UsersServiceMiddleware {
	return func(next core.UsersService) core.UsersService {
		return &usersServiceMetricsMiddleware{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			service:        next,
		}
	}
}

func (mw *usersServiceMetricsMiddleware) RegisterUser(ctx context.Context, request *domain.RegisterUserServiceRequest) (user *domain.UserDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "RegisterUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.RegisterUser(ctx, request)
}

func (mw *usersServiceMetricsMiddleware) LoginUser(ctx context.Context, request *domain.LoginUserServiceRequest) (user *domain.UserDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "LoginUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.LoginUser(ctx, request)
}

func (mw *usersServiceMetricsMiddleware) GetUser(ctx context.Context, userId int) (user *domain.UserDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "GetUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.GetUser(ctx, userId)
}

func (mw *usersServiceMetricsMiddleware) UpdateUser(ctx context.Context, request *domain.UpdateUserServiceRequest) (user *domain.UserDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "UpdateUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.UpdateUser(ctx, request)
}

func (mw *usersServiceMetricsMiddleware) GetUserProfile(ctx context.Context, username string, currentUserId int) (profile *domain.ProfileDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "GetUserProfile", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.GetUserProfile(ctx, username, currentUserId)
}

func (mw *usersServiceMetricsMiddleware) AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (profile *domain.ProfileDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "AddUserFollow", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.AddUserFollow(ctx, followerUserId, followeeUsername)
}

func (mw *usersServiceMetricsMiddleware) RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (profile *domain.ProfileDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "RemoveUserFollow", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.RemoveUserFollow(ctx, followerUserId, followeeUsername)
}
