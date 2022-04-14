package internal

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

func SeedData(logger log.Logger, usersRepository usersPersistence.UsersRepository, articlesRepository articlesPersistence.ArticlesRepository) {
	ctx := context.Background()

	var shouldSeedData bool
	{
		if _, err := usersRepository.FindUserByEmail(ctx, "user1@gmail.com"); utilities.IsNotFound(err) {
			level.Info(logger).Log("seeder", "no existing data found, continuing data seeding")
			shouldSeedData = true
		} else {
			level.Info(logger).Log("seeder", "existing data found, skipping data seeding")
			shouldSeedData = false
		}
	}

	if !shouldSeedData {
		return
	}

	// Seed some users
	user1, _ := usersRepository.CreateUser(ctx, "user1", "user1@gmail.com", "user1")
	user2, _ := usersRepository.CreateUser(ctx, "user2", "user2@gmail.com", "user2")
	user3, _ := usersRepository.CreateUser(ctx, "user3", "user3@gmail.com", "user3")

	// Seed a few articles
	_, _ = articlesRepository.CreateArticle(ctx, user1.Id, "user1 title", "user1-slug", "user1 description", "user1 body")
	_, _ = articlesRepository.CreateArticle(ctx, user1.Id, "user1 another title", "user1-another-slug", "user1 another description", "user1 another body")
	_, _ = articlesRepository.CreateArticle(ctx, user2.Id, "user2 title", "user2-slug", "user2 description", "user2 body")

	// Seed some user follows
	_, _ = usersRepository.CreateUserFollow(ctx, user2.Id, user1.Id)
	_, _ = usersRepository.CreateUserFollow(ctx, user2.Id, user3.Id)
}
