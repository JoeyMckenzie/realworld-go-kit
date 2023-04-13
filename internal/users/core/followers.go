package core

import (
	"context"
	"github.com/google/uuid"
)

func (us *userService) Follow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) error {
	return nil
}

func (us *userService) Unfollow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) error {
	return nil
}
