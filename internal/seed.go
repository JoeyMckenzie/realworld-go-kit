package internal

import (
    "context"
    "github.com/go-faker/faker/v4"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/features"
    "math/rand"
    "sync"
)

type dbSeeder struct {
    container *features.ServiceContainer
    logger    log.Logger
}

const (
    // There's a transaction limit for Planet scale, so we'll only keep 30 transactions open at a time
    usersToSeed        = 10
    randomArticleLimit = 10
)

func SeedDatabase(ctx context.Context, serviceContainer *features.ServiceContainer) {
    seeder := dbSeeder{
        container: serviceContainer,
    }

    // We'll use a wait group here as we don't need the seeding to run in sync order - fire all requests at once
    wg := new(sync.WaitGroup)

    // Keep track of all the user IDs we'll create so we can seed articles
    userIds := make([]uuid.UUID, usersToSeed)

    seeder.seedUsers(ctx, wg, &userIds)
    seeder.seedArticles(ctx, wg, &userIds)
}

func (s dbSeeder) seedUsers(ctx context.Context, wg *sync.WaitGroup, userIds *[]uuid.UUID) {

    for seedIteration := 1; seedIteration <= usersToSeed; seedIteration++ {
        wg.Add(1)
        go s.seedUser(ctx, wg, userIds)
    }

    wg.Wait()
}

func (s dbSeeder) seedUser(ctx context.Context, wg *sync.WaitGroup, userIds *[]uuid.UUID) {
    const loggingSpan = "seed_user"
    defer wg.Done()

    request := domain.AuthenticationRequest[domain.RegisterUserRequest]{
        User: &domain.RegisterUserRequest{
            //Username: fmt.Sprintf("user%d", uniqueKey),
            Email:    faker.Email(),
            Username: faker.Username(),
            // We'll reuse the same password, so we can access any unique user
            Password: "password",
        },
    }

    createdUser, err := s.container.UsersService.Register(ctx, request)

    if err != nil {
        level.Warn(s.logger).Log(loggingSpan, "error occurred while seeding a user, skipping current iteration", "err", err)
    } else {
        *userIds = append(*userIds, createdUser.ID)
    }
}

func (s dbSeeder) seedArticles(ctx context.Context, wg *sync.WaitGroup, userIds *[]uuid.UUID) {
    for _, userId := range *userIds {
        // We'll seed a random number of articles per user to simulate more active/inactive users
        randomNumberOfArticlesToSeed := rand.Intn(randomArticleLimit)
        for seedIteration := 0; seedIteration < randomNumberOfArticlesToSeed; seedIteration++ {
            wg.Add(1)
            go s.seedArticle(ctx, wg, userId)
        }
    }

    wg.Wait()
}

func (s dbSeeder) seedArticle(ctx context.Context, wg *sync.WaitGroup, userId uuid.UUID) {
    const loggingSpan = "seed_article"
    defer wg.Done()

    if userId == uuid.Nil {
        return
    }

    request := domain.CreateArticleRequest{
        Article: &domain.ArticleRequest{
            Title:       faker.Sentence(),
            Description: faker.Sentence(),
            Body:        faker.Sentence(),
            TagList: []string{
                faker.Word(),
            },
        },
    }

    if _, err := s.container.ArticlesService.CreateArticle(ctx, request, userId); err != nil {
        level.Warn(s.logger).Log(loggingSpan, "error occurred while seeding an article, skipping current iteration", "err", err)
    }
}
