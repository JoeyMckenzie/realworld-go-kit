package internal

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"sync"
	"time"
)

type dbSeeder struct {
	container *ServiceContainer
	logger    log.Logger
}

func SeedDatabase(ctx context.Context, logger log.Logger, serviceContainer *ServiceContainer) {
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

	for seedIteration := 0; seedIteration < iterations; seedIteration++ {
		go s.seedUser(ctx, wg)
	}

	wg.Wait()
}

func (s dbSeeder) seedUser(ctx context.Context, wg *sync.WaitGroup) {
	const loggingSpan = "seed_user"
	defer wg.Done()

	uniqueKey := time.Now().Nanosecond()
	request := users.AuthenticationRequest[users.RegisterUserRequest]{
		User: &users.RegisterUserRequest{
			Username: fmt.Sprintf("user%d", uniqueKey),
			Email:    fmt.Sprintf("email%d@email.com", uniqueKey),
			// We'll reuse the same password so we can access any unique user
			Password: "password",
		},
	}

	if _, err := s.container.UsersService.Register(ctx, request); err != nil {
		level.Warn(s.logger).Log(loggingSpan, "error occurred while seeding a user, skipping current iteration", "err", err)
	}
}
