package conduit_core

import (
    "github.com/go-kit/log"
    "github.com/go-playground/validator/v10"
    articlesCore "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles"
    articlesMiddlewares "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles/middlewares"
    commentsCore "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments"
    commentsMiddlewares "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/middlewares"
    tagsCore "github.com/joeymckenzie/realworld-go-kit/conduit-core/tags"
    tagsMiddlewares "github.com/joeymckenzie/realworld-go-kit/conduit-core/tags/middlewares"
    usersCore "github.com/joeymckenzie/realworld-go-kit/conduit-core/users"
    usersMiddlewares "github.com/joeymckenzie/realworld-go-kit/conduit-core/users/middlewares"
    "github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
)

// ConduitServiceRegister houses all application services used within the core layer. Each service comprises a primary
// business logic service wrapped by various middleware orchestrated by go-kit to facilitate logging, validation, etc.
type ConduitServiceRegister struct {
    UsersService    usersCore.UsersService
    ArticlesService articlesCore.ArticlesService
    CommentsService commentsCore.CommentsService
    TagsService     tagsCore.TagsService
}

func newConduitServiceRegister(usersService usersCore.UsersService, articlesService articlesCore.ArticlesService, commentsService commentsCore.CommentsService, tagsService tagsCore.TagsService) *ConduitServiceRegister {
    return &ConduitServiceRegister{
        UsersService:    usersService,
        ArticlesService: articlesService,
        CommentsService: commentsService,
        TagsService:     tagsService,
    }
}

func InitializeServices(logger log.Logger, entClient *ent.Client) *ConduitServiceRegister {
    requestValidator := validator.New()

    var usersService usersCore.UsersService
    {
        requestCount, requestLatency := utilities.NewServiceMetrics("users_service")
        usersService = usersCore.NewUsersService(requestValidator, entClient, services.NewTokenService(), services.NewSecurityService())
        usersService = usersMiddlewares.NewUsersServiceLoggingMiddleware(logger)(usersService)
        usersService = usersMiddlewares.NewUsersServiceMetricsMiddleware(requestCount, requestLatency)(usersService)
        usersService = usersMiddlewares.NewUsersServiceRequestValidationMiddleware(logger, requestValidator)(usersService)
    }

    var articlesService articlesCore.ArticlesService
    {
        requestCount, requestLatency := utilities.NewServiceMetrics("articles_service")
        articlesService = articlesCore.NewArticlesServices(requestValidator, entClient)
        articlesService = articlesMiddlewares.NewArticlesServiceLoggingMiddleware(logger)(articlesService)
        articlesService = articlesMiddlewares.NewArticlesServiceMetricsMiddleware(requestCount, requestLatency)(articlesService)
        articlesService = articlesMiddlewares.NewArticlesServiceRequestValidationMiddleware(logger, requestValidator)(articlesService)
    }

    var commentsService commentsCore.CommentsService
    {
        requestCount, requestLatency := utilities.NewServiceMetrics("comments_service")
        commentsService = commentsCore.NewCommentsService(requestValidator, entClient)
        commentsService = commentsMiddlewares.NewCommentsServiceLoggingMiddleware(logger)(commentsService)
        commentsService = commentsMiddlewares.NewCommentsServiceMetricsMiddleware(requestCount, requestLatency)(commentsService)
        commentsService = commentsMiddlewares.NewCommentsServiceRequestValidationMiddleware(logger, requestValidator)(commentsService)
    }

    var tagsService tagsCore.TagsService
    {
        requestCount, requestLatency := utilities.NewServiceMetrics("tags_service")
        tagsService = tagsCore.NewTagsService(requestValidator, entClient)
        tagsService = tagsMiddlewares.NewTagsServiceLoggingMiddleware(logger)(tagsService)
        tagsService = tagsMiddlewares.NewTagsServiceMetricsMiddleware(requestCount, requestLatency)(tagsService)
    }

    return newConduitServiceRegister(usersService, articlesService, commentsService, tagsService)
}
