package core

import (
    "context"
    "database/sql"
    "github.com/go-playground/validator/v10"
    "github.com/gosimple/slug"
    "github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
    usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
    "github.com/joeymckenzie/realworld-go-kit/pkg/api"
    "github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
    "net/http"
)

type (
    ArticlesService interface {
        GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) (*[]domain.ArticleDto, error)
        CreateArticle(ctx context.Context, request *domain.UpsertArticleServiceRequest) (*domain.ArticleDto, error)
    }

    articlesServices struct {
        validator       *validator.Validate
        repository      persistence.ArticlesRepository
        usersRepository usersPersistence.UsersRepository
    }

    ArticlesServiceMiddleware func(articlesService ArticlesService) ArticlesService
)

func NewArticlesServices(validator *validator.Validate, repository persistence.ArticlesRepository, usersRepository usersPersistence.UsersRepository) ArticlesService {
    return &articlesServices{
        validator:       validator,
        repository:      repository,
        usersRepository: usersRepository,
    }
}

func (as *articlesServices) GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) (*[]domain.ArticleDto, error) {
    articles, err := as.repository.GetArticles(ctx, request)
    articlesNotFound := articles == nil || utilities.IsNotFound(err) || len(*articles) == 0

    if utilities.IsValidDbError(err) {
        return nil, api.NewInternalServerErrorWithContext("articles", err)
    } else if articlesNotFound {
        return nil, api.NewApiErrorWithContext(http.StatusNotFound, "articles", utilities.ErrArticlesNotFound)
    }

    var response []domain.ArticleDto

    for _, article := range *articles {
        // Get the user that authored the article
        user, err := as.usersRepository.GetUser(ctx, article.UserId)
        if err != nil {
            return nil, api.NewInternalServerErrorWithContext("articles", err)
        }

        // Get any associated tags from the article
        articleTags, err := as.repository.GetArticleTags(ctx, article.Id)
        if utilities.IsValidDbError(err) {
            return nil, api.NewInternalServerErrorWithContext("articles", err)
        }

        // Map the article tags, if any were found
        var mappedArticleTags []string
        if articleTags != nil {
            for _, tag := range *articleTags {
                mappedArticleTags = append(mappedArticleTags, tag)
            }
        }

        // If the user ID is included on the request, check if they've favorited the article or followed the author
        // favoritedArticle := false
        followingUser := false
        if request.UserId > -1 {
            followingUser = as.usersRepository.IsFollowingUser(ctx, request.UserId, article.UserId)
        }

        mappedArticle := &domain.ArticleDto{
            Slug:           article.Slug,
            Title:          article.Title,
            Description:    article.Description,
            Body:           article.Body,
            TagList:        mappedArticleTags,
            CreatedAt:      article.CreatedAt,
            UpdatedAt:      article.UpdatedAt,
            Favorited:      false,
            FavoritesCount: article.Favorites,
            Author: domain.AuthorDto{
                Username:  user.Username,
                Bio:       user.Bio,
                Image:     user.Image,
                Following: followingUser,
            },
        }

        response = append(response, *mappedArticle)
    }

    return &response, nil
}

func (as *articlesServices) CreateArticle(ctx context.Context, request *domain.UpsertArticleServiceRequest) (*domain.ArticleDto, error) {
    // Verify the user exists, ensure no article is created without a valid existing user
    existingUser, err := as.usersRepository.GetUser(ctx, request.UserId)
    if err != nil || existingUser == nil {
        return nil, api.NewApiErrorWithContext(http.StatusConflict, "articles", utilities.ErrUserNotFound)
    }

    // Verify the article title slug is unique
    articleSlug := slug.Make(request.Title)
    existingArticleFromSlug, err := as.repository.FindArticleBySlug(ctx, articleSlug)

    if err != nil && err != sql.ErrNoRows {
        return nil, err
    } else if existingArticleFromSlug != nil {
        return nil, api.NewApiErrorWithContext(http.StatusConflict, "article", utilities.ErrArticleTitleExists)
    }

    // Dedupe the request list of tags, initializing an empty list if none is provided on the request
    tagsToCreate := removeDuplicates(request.TagList)

    var articleTagsToCreate []int
    {
        // Create any tags on the request, if supplied
        if len(tagsToCreate) > 0 {
            // Get the existing tags for checking against those on the request
            existingTags, err := as.repository.GetTags(ctx, tagsToCreate)
            if err != nil && err != sql.ErrNoRows {
                return nil, api.NewInternalServerErrorWithContext("tags", err)
            }

            // Roll through the tags on the request to see if we should create any new tags
            for _, tag := range tagsToCreate {
                // If the tag already exists, skip creating it and add it to the list of reference IDs for the article
                if existingTag := findTag(tag, existingTags); existingTag != nil {
                    articleTagsToCreate = append(articleTagsToCreate, existingTag.Id)
                    continue
                }

                // We've detected a new tag at this point, create it and rollup any errors
                createdTag, err := as.repository.CreateTag(ctx, tag)
                if err != nil {
                    return nil, api.NewInternalServerErrorWithContext("tags", err)
                }

                // Add the newly created tag ID to the list to reference from articles
                articleTagsToCreate = append(articleTagsToCreate, createdTag.Id)
            }
        }
    }

    // Create the article, adding the article tags as a text array type in Postgres
    createdArticle, err := as.repository.CreateArticle(ctx, request.UserId, request.Title, articleSlug, request.Description, request.Body)
    if err != nil {
        return nil, api.NewInternalServerErrorWithContext("articles", err)
    }

    // Create the associated article tags
    for _, tagId := range articleTagsToCreate {
        if _, err := as.repository.CreateArticleTag(ctx, tagId, createdArticle.Id); err != nil {
            return nil, api.NewInternalServerErrorWithContext("articleTags", err)
        }
    }

    return &domain.ArticleDto{
        Slug:           articleSlug,
        Title:          createdArticle.Title,
        Description:    createdArticle.Description,
        Body:           createdArticle.Body,
        TagList:        tagsToCreate,
        CreatedAt:      createdArticle.CreatedAt,
        UpdatedAt:      createdArticle.UpdatedAt,
        Favorited:      false,
        FavoritesCount: 0,
        Author: domain.AuthorDto{
            Username:  existingUser.Username,
            Bio:       existingUser.Bio,
            Image:     existingUser.Image,
            Following: false,
        },
    }, nil
}
