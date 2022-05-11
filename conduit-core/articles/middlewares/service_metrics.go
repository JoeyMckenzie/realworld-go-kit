package middlewares

import (
    "context"
    "fmt"
    "github.com/go-kit/kit/metrics"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles"
    articlesDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/articles"
    "time"
)

type articlesServiceMetricsMiddleware struct {
    requestCount   metrics.Counter
    requestLatency metrics.Histogram
    service        articles.ArticlesService
}

func NewArticlesServiceMetricsMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) articles.ArticlesServiceMiddleware {
    return func(next articles.ArticlesService) articles.ArticlesService {
        return &articlesServiceMetricsMiddleware{
            requestCount:   requestCount,
            requestLatency: requestLatency,
            service:        next,
        }
    }
}

func (mw *articlesServiceMetricsMiddleware) GetArticles(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) (articles []*articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "GetArticles", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.GetArticles(ctx, request)
}

func (mw *articlesServiceMetricsMiddleware) GetArticle(ctx context.Context, request *articlesDomain.GetArticleServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "GetArticle", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.GetArticle(ctx, request)
}

func (mw *articlesServiceMetricsMiddleware) CreateArticle(ctx context.Context, request *articlesDomain.CreateArticleServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "CreateArticle", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.CreateArticle(ctx, request)
}

func (mw *articlesServiceMetricsMiddleware) GetFeed(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) (articles []*articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "GetFeed", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.GetFeed(ctx, request)
}

func (mw *articlesServiceMetricsMiddleware) UpdateArticle(ctx context.Context, request *articlesDomain.UpdateArticleServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "UpdateArticle", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.UpdateArticle(ctx, request)
}

func (mw *articlesServiceMetricsMiddleware) DeleteArticle(ctx context.Context, request *articlesDomain.DeleteArticleServiceRequest) (err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "DeleteArticle", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.DeleteArticle(ctx, request)
}

func (mw *articlesServiceMetricsMiddleware) FavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "FavoriteArticle", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.FavoriteArticle(ctx, request)
}

func (mw *articlesServiceMetricsMiddleware) UnfavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "UnfavoriteArticle", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.UnfavoriteArticle(ctx, request)
}
