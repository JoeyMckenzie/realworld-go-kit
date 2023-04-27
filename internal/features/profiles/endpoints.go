package profiles

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func makeGetProfileEndpoint(service ProfilesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        usernameToFollow := ctx.Value(shared.UsernameContextKey{}).(shared.UsernameContextKey)
        profile, err := service.GetProfile(ctx, usernameToFollow.Username, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ProfileResponse{
            Profile: profile,
        }, nil
    }
}

func makeFollowUserEndpoint(service ProfilesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        usernameToFollow := ctx.Value(shared.UsernameContextKey{}).(shared.UsernameContextKey)
        profile, err := service.Follow(ctx, usernameToFollow.Username, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ProfileResponse{
            Profile: profile,
        }, nil
    }
}

func makeUnfollowUserEndpoint(service ProfilesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        usernameToUnfollow := ctx.Value(shared.UsernameContextKey{}).(shared.UsernameContextKey)
        profile, err := service.Unfollow(ctx, usernameToUnfollow.Username, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ProfileResponse{
            Profile: profile,
        }, nil
    }
}
