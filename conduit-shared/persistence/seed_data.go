package persistence

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
)

// SeedData uses ent to seed the database with test data for users and articles
func SeedData(ctx context.Context, client *ent.Client) {
	// Check for existing users, bail out of seeding if users already exist
	_, err := client.User.
		Query().
		First(ctx)

	if !ent.IsNotFound(err) {
		return
	}

	users := seedUsers(ctx, client)
	articles := seedArticles(ctx, client, users)
	tags := seedTags(ctx, client)
	seedFollows(ctx, client, users...)
	seedFavorites(ctx, client, users, articles)
	seedArticleTags(ctx, client, articles, tags)
	seedComments(ctx, client, users, articles)
}

func seedUsers(ctx context.Context, client *ent.Client) []*ent.User {
	securityService := services.NewSecurityService()
	testPassword1, _ := securityService.HashPassword("testPassword1")
	testPassword2, _ := securityService.HashPassword("testPassword2")
	testPassword3, _ := securityService.HashPassword("testPassword3")

	users := []*ent.User{
		client.User.
			Create().
			SetUsername("testUser1").
			SetEmail("testUser1@gmail.com").
			SetPassword(testPassword1).
			SetBio("testUser1 bio").
			SetImage("testUser1 image").
			SaveX(ctx),
		client.User.
			Create().
			SetUsername("testUser2").
			SetEmail("testUser2@gmail.com").
			SetPassword(testPassword2).
			SetBio("testUser2 bio").
			SetImage("testUser2 image").
			SaveX(ctx),
		client.User.
			Create().
			SetUsername("testUser3").
			SetEmail("testUser3@gmail.com").
			SetPassword(testPassword3).
			SetBio("testUser3 bio").
			SetImage("testUser3 image").
			SaveX(ctx),
	}

	return users
}

func seedFollows(ctx context.Context, client *ent.Client, users ...*ent.User) {
	// testUser2 follows testUser1
	client.Follow.
		Create().
		SetUserFollower(users[1]).
		SetUserFollowee(users[0]).
		SaveX(ctx)

	// testUser3 follows testUser1
	client.Follow.
		Create().
		SetUserFollower(users[2]).
		SetUserFollowee(users[0]).
		SaveX(ctx)

	// testUser3 follows testUser2
	client.Follow.
		Create().
		SetUserFollower(users[2]).
		SetUserFollowee(users[1]).
		SaveX(ctx)
}

func seedTags(ctx context.Context, client *ent.Client) []*ent.Tag {
	tags := []*ent.Tag{
		client.Tag.
			Create().
			SetTag("tag1").
			SaveX(ctx),
		client.Tag.
			Create().
			SetTag("tag2").
			SaveX(ctx),
		client.Tag.
			Create().
			SetTag("tag3").
			SaveX(ctx),
	}

	return tags
}

func seedArticleTags(ctx context.Context, client *ent.Client, articles []*ent.Article, tags []*ent.Tag) {
	// Add "tag1" to "testUser1 article"
	client.ArticleTag.
		Create().
		SetTag(tags[0]).
		SetArticle(articles[0]).
		SaveX(ctx)

	// Add "tag2" to "testUser1 article"
	client.ArticleTag.
		Create().
		SetTag(tags[1]).
		SetArticle(articles[0]).
		SaveX(ctx)

	// Add "tag1" to "testUser1 another article"
	client.ArticleTag.
		Create().
		SetTag(tags[1]).
		SetArticle(articles[1]).
		SaveX(ctx)

	// Add "tag3" to "testUser2 article"
	client.ArticleTag.
		Create().
		SetTag(tags[2]).
		SetArticle(articles[2]).
		SaveX(ctx)
}

func seedArticles(ctx context.Context, client *ent.Client, users []*ent.User) []*ent.Article {
	articlesToCreate := []*ent.ArticleCreate{
		client.Article.
			Create().
			SetSlug(slug.Make("testUser1 article")).
			SetTitle("testUser1 article").
			SetDescription("testUser1 description").
			SetBody("testUser1 body").
			SetAuthor(users[0]),
		client.Article.
			Create().
			SetSlug(slug.Make("testUser1 another article")).
			SetTitle("testUser1 another article").
			SetDescription("testUser1 another description").
			SetBody("testUser1 another body").
			SetAuthor(users[0]),
		client.Article.
			Create().
			SetSlug(slug.Make("testUser2 article")).
			SetTitle("testUser2 article").
			SetDescription("testUser2 description").
			SetBody("testUser2 body").
			SetAuthor(users[1]),
	}

	createdArticles := client.Article.
		CreateBulk(articlesToCreate...).
		SaveX(ctx)

	return createdArticles
}

func seedFavorites(ctx context.Context, client *ent.Client, users []*ent.User, articles []*ent.Article) {
	// testUser2 favorites "testUser1 article"
	client.Favorite.Create().
		SetArticleFavorites(articles[0]).
		SetUserFavorites(users[1]).
		SaveX(ctx)

	// testUser2 favorites "testUser1 another article"
	client.Favorite.Create().
		SetArticleFavorites(articles[1]).
		SetUserFavorites(users[1]).
		SaveX(ctx)

	// testUser3 favorites "testUser1 article"
	client.Favorite.Create().
		SetArticleFavorites(articles[0]).
		SetUserFavorites(users[2]).
		SaveX(ctx)

	// testUser3 favorites "testUser2 article"
	client.Favorite.Create().
		SetArticleFavorites(articles[1]).
		SetUserFavorites(users[2]).
		SaveX(ctx)
}

func seedComments(ctx context.Context, client *ent.Client, users []*ent.User, articles []*ent.Article) []*ent.Comment {
	commentsToCreate := []*ent.CommentCreate{
		// testUser2 adds comment to testuser1-article
		client.Comment.
			Create().
			SetBody("testUser2 comment").
			SetUserID(users[1].ID).
			SetArticleID(articles[0].ID),
		// testUser3 adds comment to testuser1-article
		client.Comment.
			Create().
			SetBody("testUser3 comment").
			SetUserID(users[2].ID).
			SetArticleID(articles[0].ID),
		// testUser1 adds comment to testuser2-article
		client.Comment.
			Create().
			SetBody("testUser1 comment").
			SetUserID(users[0].ID).
			SetArticleID(articles[2].ID),
	}

	createdComments := client.Comment.
		CreateBulk(commentsToCreate...).
		SaveX(ctx)

	return createdComments
}
