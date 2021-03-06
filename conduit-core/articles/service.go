package articles

import (
    "context"
    "github.com/go-playground/validator/v10"
    "github.com/gosimple/slug"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
    articlesDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/articles"
    sharedDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/shared"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/article"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/articletag"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/favorite"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/tag"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/user"
    sharedUtilities "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "net/http"
    "time"
)

var DefaultArticlesResponse = make([]*articlesDomain.ArticleDto, 0)

type (
    ConduitArticlesService interface {
        GetArticles(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) ([]*articlesDomain.ArticleDto, error)
        GetArticle(ctx context.Context, request *articlesDomain.GetArticleServiceRequest) (*articlesDomain.ArticleDto, error)
        GetFeed(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) ([]*articlesDomain.ArticleDto, error)
        CreateArticle(ctx context.Context, request *articlesDomain.CreateArticleServiceRequest) (*articlesDomain.ArticleDto, error)
        UpdateArticle(ctx context.Context, request *articlesDomain.UpdateArticleServiceRequest) (*articlesDomain.ArticleDto, error)
        DeleteArticle(ctx context.Context, request *articlesDomain.DeleteArticleServiceRequest) error
        FavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (*articlesDomain.ArticleDto, error)
        UnfavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (*articlesDomain.ArticleDto, error)
    }

    articlesService struct {
        validator *validator.Validate
        client    *ent.Client
    }

    ConduitArticlesServiceMiddleware func(articlesService ConduitArticlesService) ConduitArticlesService
)

func NewArticlesServices(validator *validator.Validate, client *ent.Client) ConduitArticlesService {
    return &articlesService{
        validator: validator,
        client:    client,
    }
}

func (as *articlesService) GetArticles(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) ([]*articlesDomain.ArticleDto, error) {
    articles := as.client.Article.
        Query().
        WithFavorites().
        WithAuthor(func(query *ent.UserQuery) {
            query.WithFollowees()
        }).
        WithArticleTags(func(query *ent.ArticleTagQuery) {
            query.WithTag()
        })

    if request.Tag != "" {
        // Check for the existing tag on the request
        tagId, err := as.client.Tag.
            Query().
            Where(tag.Tag(request.Tag)).
            Select(tag.FieldID).
            Int(ctx)

        if ent.IsNotFound(err) {
            // Tag doesn't exist, so no articles can have the associated tag
            return DefaultArticlesResponse, nil
        }

        articles.Where(article.HasArticleTagsWith(articletag.TagID(tagId)))
    }

    if request.Author != "" {
        articles.Where(article.HasAuthorWith(user.Username(request.Author)))
    }

    if request.Favorited != "" {
        // Verify the favoriting user actually exists
        favoritingUserId, err := as.client.User.
            Query().
            Where(user.Username(request.Favorited)).
            Select(user.FieldID).
            Int(ctx)

        if ent.IsNotFound(err) {
            // User doesn't exist, so no articles are favorited by them
            return DefaultArticlesResponse, nil
        }

        articles.Where(article.HasFavoritesWith(favorite.UserID(favoritingUserId)))
    }

    queriedArticles, err := articles.
        Offset(request.Offset).
        Limit(request.Limit).
        Order(ent.Desc(article.FieldCreateTime)).
        All(ctx)

    return makeArticlesMapping(queriedArticles, request, false, err)
}

func (as *articlesService) GetFeed(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) ([]*articlesDomain.ArticleDto, error) {
    articles, err := as.client.Article.
        Query().
        WithFavorites().
        WithAuthor().
        WithArticleTags(func(query *ent.ArticleTagQuery) {
            query.WithTag()
        }).
        Where(article.HasFavoritesWith(favorite.UserID(request.UserId))).
        Offset(request.Offset).
        Limit(request.Limit).
        Order(ent.Desc(article.FieldCreateTime)).
        All(ctx)

    return makeArticlesMapping(articles, request, true, err)
}

func (as *articlesService) GetArticle(ctx context.Context, request *articlesDomain.GetArticleServiceRequest) (*articlesDomain.ArticleDto, error) {
    existingArticle, err := as.client.Article.
        Query().
        WithFavorites().
        WithAuthor().
        WithArticleTags(func(query *ent.ArticleTagQuery) {
            query.WithTag()
        }).
        Where(article.Slug(request.Slug)).
        First(ctx)

    if ent.IsNotFound(err) {
        return nil, shared.NewApiErrorWithContext(http.StatusNotFound, "article", sharedUtilities.ErrArticlesNotFound)
    }

    return makeArticleMapping(existingArticle, false, request.UserId), nil
}

func (as *articlesService) CreateArticle(ctx context.Context, request *articlesDomain.CreateArticleServiceRequest) (*articlesDomain.ArticleDto, error) {
    // Verify the user exists, ensure no article is created without a valid existing user
    existingUser, err := as.client.User.Get(ctx, request.UserId)

    if ent.IsNotFound(err) {
        return nil, shared.NewApiErrorWithContext(http.StatusBadRequest, "articles", sharedUtilities.ErrUserNotFound)
    }

    // Verify the article title slug is unique
    articleSlug := slug.Make(request.Title)
    _, err = as.client.Article.
        Query().
        Where(article.Slug(articleSlug)).
        First(ctx)

    if !ent.IsNotFound(err) {
        return nil, shared.NewApiErrorWithContext(http.StatusConflict, "article", sharedUtilities.ErrArticleTitleExists)
    }

    // Dedupe the request list of tags, initializing an empty list if none is provided on the request
    tagsToCreate := removeDuplicates(request.TagList)

    // Create a transaction so that we can rollback any inserts on error
    transaction, err := as.client.BeginTx(ctx, nil)

    if err != nil {
        return nil, shared.NewInternalServerErrorWithContext("tags", err)
    }

    articleTagsToBulkInsert, err := as.makeArticleTagsMapping(ctx, transaction, tagsToCreate)

    // Error mapping is handled while mapping article tags, propagate it up
    if err != nil {
        return nil, err
    }

    articleToCreate := transaction.Article.Create().
        SetAuthorID(request.UserId).
        SetTitle(request.Title).
        SetBody(request.Body).
        SetSlug(articleSlug).
        SetDescription(request.Description)

    if len(articleTagsToBulkInsert) > 0 {
        createdArticleTags, err := transaction.ArticleTag.
            CreateBulk(articleTagsToBulkInsert...).
            Save(ctx)

        if err != nil {
            _ = transaction.Rollback()
            return nil, shared.NewInternalServerErrorWithContext("article_tags", err)
        }

        articleToCreate.AddArticleTags(createdArticleTags...)
    }

    // Create the article, adding the article tags as a text array type in Postgres
    createdArticle, err := articleToCreate.Save(ctx)

    if err != nil {
        _ = transaction.Rollback()
        return nil, shared.NewInternalServerErrorWithContext("articles", err)
    }

    // Finally, commit the transaction
    if err = transaction.Commit(); err != nil {
        _ = transaction.Rollback()
        return nil, shared.NewInternalServerErrorWithContext("articles", err)
    }

    return &articlesDomain.ArticleDto{
        Slug:           createdArticle.Slug,
        Title:          createdArticle.Title,
        Description:    createdArticle.Description,
        Body:           createdArticle.Body,
        TagList:        tagsToCreate,
        CreatedAt:      createdArticle.CreateTime,
        UpdatedAt:      createdArticle.UpdateTime,
        Favorited:      false,
        FavoritesCount: 0,
        Author: sharedDomain.AuthorDto{
            Username:  existingUser.Username,
            Bio:       existingUser.Bio,
            Image:     existingUser.Image,
            Following: false,
        },
    }, nil
}

func (as *articlesService) UpdateArticle(ctx context.Context, request *articlesDomain.UpdateArticleServiceRequest) (*articlesDomain.ArticleDto, error) {
    existingArticle, err := as.client.Article.
        Query().
        Where(
            article.Slug(request.ArticleSlug),
            article.UserID(request.UserId),
        ).
        WithFavorites().
        WithAuthor().
        WithArticleTags(func(query *ent.ArticleTagQuery) {
            query.WithTag()
        }).
        First(ctx)

    // If the article is not found, i.e. user-article ID mismatch, do not let malicious users know
    // a resource they do not own, that requires authentication to modify, exists. Obfuscate the
    // response by simply returning a not found rather than forbidden for better security experience.
    if ent.IsNotFound(err) {
        return nil, shared.NewApiErrorWithContext(http.StatusNotFound, "article", sharedUtilities.ErrArticlesNotFound)
    }

    updatedTitle := sharedUtilities.UpdateIfRequired(existingArticle.Title, request.Title)
    updatedDescription := sharedUtilities.UpdateIfRequired(existingArticle.Description, request.Description)
    updatedBody := sharedUtilities.UpdateIfRequired(existingArticle.Title, request.Body)
    updatedSlug := slug.Make(updatedTitle)

    if updatedTitle != existingArticle.Title {
        // Verify the updated slug title is available
        existingSlug, _ := as.client.Article.
            Query().
            Where(article.Slug(updatedSlug)).
            First(ctx)

        if existingSlug != nil {
            return nil, shared.NewApiErrorWithContext(http.StatusConflict, "article", sharedUtilities.ErrArticleTitleExists)
        }
    }

    updatedArticle, err := as.client.Article.
        UpdateOne(existingArticle).
        SetTitle(updatedTitle).
        SetSlug(updatedSlug).
        SetDescription(updatedDescription).
        SetBody(updatedBody).
        SetUpdateTime(time.Now()).
        Save(ctx)

    if err != nil {
        return nil, shared.NewInternalServerErrorWithContext("article", err)
    }

    var tagList []string
    {
        for _, existingTag := range existingArticle.Edges.ArticleTags {
            if existingTag.Edges.Tag != nil {
                tagList = append(tagList, existingTag.Edges.Tag.Tag)
            }
        }
    }

    return &articlesDomain.ArticleDto{
        Slug:           updatedArticle.Slug,
        Title:          updatedArticle.Title,
        Description:    updatedArticle.Description,
        Body:           updatedArticle.Body,
        TagList:        tagList,
        CreatedAt:      updatedArticle.CreateTime,
        UpdatedAt:      updatedArticle.UpdateTime,
        Favorited:      false,
        FavoritesCount: len(existingArticle.Edges.Favorites),
        Author: sharedDomain.AuthorDto{
            Username:  existingArticle.Edges.Author.Username,
            Bio:       existingArticle.Edges.Author.Bio,
            Image:     existingArticle.Edges.Author.Image,
            Following: false,
        },
    }, nil
}

func (as *articlesService) DeleteArticle(ctx context.Context, request *articlesDomain.DeleteArticleServiceRequest) error {
    existingArticle, err := as.client.Article.
        Query().
        Where(
            article.Slug(request.ArticleSlug),
            article.UserID(request.UserId),
        ).
        First(ctx)

    // If the article is not found, i.e. user-article ID mismatch, do not let malicious users know
    // a resource they do not own, that requires authentication to modify, exists. Obfuscate the
    // response by simply returning a not found rather than forbidden for better security experience.
    if ent.IsNotFound(err) {
        return shared.NewApiErrorWithContext(http.StatusNotFound, "article", sharedUtilities.ErrArticlesNotFound)
    }

    err = as.client.Article.
        DeleteOne(existingArticle).
        Exec(ctx)

    if err != nil {
        return shared.NewInternalServerErrorWithContext("article", err)
    }

    return nil
}

func (as *articlesService) FavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (*articlesDomain.ArticleDto, error) {
    existingArticle, err := as.getExistingArticleForFavoriting(ctx, request)

    // Error is converted in our utility method, so pass it back up the stack
    if err != nil {
        return nil, err
    }

    _, err = as.client.Favorite.
        Create().
        SetArticleID(existingArticle.ID).
        SetUserID(request.UserId).
        Save(ctx)

    if err != nil {
        return nil, shared.NewInternalServerErrorWithContext("article", err)
    }

    // TODO: Not great, may be better to update the already existing entity in-mem rather than re-retrieving
    existingArticle, err = as.getExistingArticleForFavoriting(ctx, request)

    // Error is converted in our utility method, so pass it back up the stack
    if err != nil {
        return nil, err
    }

    return makeArticleMapping(existingArticle, true, request.UserId), nil
}

func (as *articlesService) UnfavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (*articlesDomain.ArticleDto, error) {
    existingArticle, err := as.getExistingArticleForFavoriting(ctx, request)

    // Error is converted in our utility method, so pass it back up the stack
    if err != nil {
        return nil, err
    }

    if _, err = as.client.Favorite.
        Delete().
        Where(
            favorite.ArticleID(existingArticle.ID),
            favorite.UserID(request.UserId),
        ).
        Exec(ctx); err != nil {
        return nil, shared.NewInternalServerErrorWithContext("article_favorite", err)
    }

    // TODO: Not great, may be better to update the already existing entity in-mem rather than re-retrieving
    existingArticle, err = as.getExistingArticleForFavoriting(ctx, request)

    // Error is converted in our utility method, so pass it back up the stack
    if err != nil {
        return nil, err
    }

    return makeArticleMapping(existingArticle, false, request.UserId), nil
}
