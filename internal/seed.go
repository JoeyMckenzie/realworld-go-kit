package internal

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/features"
	"sync"
	"time"
)

type dbSeeder struct {
	container *features.ServiceContainer
	logger    log.Logger
}

func SeedDatabase(ctx context.Context, serviceContainer *features.ServiceContainer) {
	seeder := dbSeeder{
		container: serviceContainer,
	}

	seeder.seedUsers(ctx)
}

func (s dbSeeder) seedUsers(ctx context.Context) {
	const iterations = 100

	// We'll use a wait group here as we don't need the seeding to run in sync order - fire all requests at once
	wg := new(sync.WaitGroup)
	wg.Add(iterations)

	// We'll take every tenth user created, and from there generate 10 articles each
	userIds := make([]uuid.UUID, 10)

	for seedIteration := 1; seedIteration <= iterations; seedIteration++ {
		go s.seedUser(ctx, wg, &userIds, seedIteration%10 == 0)
	}

	wg.Wait()
}

func (s dbSeeder) seedUser(ctx context.Context, wg *sync.WaitGroup, userIds *[]uuid.UUID, addUserToArticlesSeed bool) {
	const loggingSpan = "seed_user"
	defer wg.Done()

	uniqueKey := time.Now().Nanosecond()
	request := domain.AuthenticationRequest[domain.RegisterUserRequest]{
		User: &domain.RegisterUserRequest{
			Username: fmt.Sprintf("user%d", uniqueKey),
			Email:    fmt.Sprintf("email%d@email.com", uniqueKey),
			// We'll reuse the same password, so we can access any unique user
			Password: "password",
		},
	}

	createdUser, err := s.container.UsersService.Register(ctx, request)

	if err != nil {
		level.Warn(s.logger).Log(loggingSpan, "error occurred while seeding a user, skipping current iteration", "err", err)
	}

	if addUserToArticlesSeed {
		*userIds = append(*userIds, createdUser.ID)
	}
}
