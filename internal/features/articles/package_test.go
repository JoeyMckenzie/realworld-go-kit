package articles

import (
    "context"
    "github.com/go-faker/faker/v4"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/stretchr/testify/suite"
    "golang.org/x/exp/slog"
    "os"
    "testing"

    _ "github.com/go-sql-driver/mysql"
)

type ArticlesServiceTestSuite struct {
    suite.Suite
    Ctx        context.Context
    Service    ArticlesService
    SeedUserId uuid.UUID
    Db         *sqlx.DB
}

func (s *ArticlesServiceTestSuite) SetupSuite() {
    // Setup our atomic dependencies
    ctx := context.Background()
    nopLogger := slog.Default()
    dsn := os.Getenv("DSN")
    db := sqlx.MustOpen("mysql", dsn)

    // Next, setup our service dependencies
    articlesRepository := repositories.NewArticlesRepository(db)
    usersRepository := repositories.NewUsersRepository(db)

    // Seed an existing user to use for assertions
    seedUser, _ := usersRepository.CreateUser(ctx, faker.Username(), faker.Email(), faker.PASSWORD)

    s.Ctx = ctx
    s.Service = NewArticlesService(nopLogger, articlesRepository, usersRepository)
    s.SeedUserId = seedUser.ID
}

func Test_RunArticlesTestSuite(t *testing.T) {
    suite.Run(t, new(ArticlesServiceTestSuite))
}
