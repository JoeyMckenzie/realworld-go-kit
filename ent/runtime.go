// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/joeymckenzie/realworld-go-kit/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/ent/articletag"
	"github.com/joeymckenzie/realworld-go-kit/ent/favorite"
	"github.com/joeymckenzie/realworld-go-kit/ent/follow"
	"github.com/joeymckenzie/realworld-go-kit/ent/schema"
	"github.com/joeymckenzie/realworld-go-kit/ent/tag"
	"github.com/joeymckenzie/realworld-go-kit/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleMixin := schema.Article{}.Mixin()
	articleMixinFields0 := articleMixin[0].Fields()
	_ = articleMixinFields0
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescCreateTime is the schema descriptor for create_time field.
	articleDescCreateTime := articleMixinFields0[0].Descriptor()
	// article.DefaultCreateTime holds the default value on creation for the create_time field.
	article.DefaultCreateTime = articleDescCreateTime.Default.(func() time.Time)
	// articleDescUpdateTime is the schema descriptor for update_time field.
	articleDescUpdateTime := articleMixinFields0[1].Descriptor()
	// article.DefaultUpdateTime holds the default value on creation for the update_time field.
	article.DefaultUpdateTime = articleDescUpdateTime.Default.(func() time.Time)
	// article.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	article.UpdateDefaultUpdateTime = articleDescUpdateTime.UpdateDefault.(func() time.Time)
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[0].Descriptor()
	// article.DefaultTitle holds the default value on creation for the title field.
	article.DefaultTitle = articleDescTitle.Default.(string)
	// article.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	article.TitleValidator = articleDescTitle.Validators[0].(func(string) error)
	// articleDescBody is the schema descriptor for body field.
	articleDescBody := articleFields[1].Descriptor()
	// article.DefaultBody holds the default value on creation for the body field.
	article.DefaultBody = articleDescBody.Default.(string)
	// article.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	article.BodyValidator = articleDescBody.Validators[0].(func(string) error)
	// articleDescDescription is the schema descriptor for description field.
	articleDescDescription := articleFields[2].Descriptor()
	// article.DefaultDescription holds the default value on creation for the description field.
	article.DefaultDescription = articleDescDescription.Default.(string)
	// article.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	article.DescriptionValidator = articleDescDescription.Validators[0].(func(string) error)
	// articleDescSlug is the schema descriptor for slug field.
	articleDescSlug := articleFields[3].Descriptor()
	// article.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	article.SlugValidator = articleDescSlug.Validators[0].(func(string) error)
	articletagMixin := schema.ArticleTag{}.Mixin()
	articletagMixinFields0 := articletagMixin[0].Fields()
	_ = articletagMixinFields0
	articletagFields := schema.ArticleTag{}.Fields()
	_ = articletagFields
	// articletagDescCreateTime is the schema descriptor for create_time field.
	articletagDescCreateTime := articletagMixinFields0[0].Descriptor()
	// articletag.DefaultCreateTime holds the default value on creation for the create_time field.
	articletag.DefaultCreateTime = articletagDescCreateTime.Default.(func() time.Time)
	favoriteMixin := schema.Favorite{}.Mixin()
	favoriteMixinFields0 := favoriteMixin[0].Fields()
	_ = favoriteMixinFields0
	favoriteFields := schema.Favorite{}.Fields()
	_ = favoriteFields
	// favoriteDescCreateTime is the schema descriptor for create_time field.
	favoriteDescCreateTime := favoriteMixinFields0[0].Descriptor()
	// favorite.DefaultCreateTime holds the default value on creation for the create_time field.
	favorite.DefaultCreateTime = favoriteDescCreateTime.Default.(func() time.Time)
	followMixin := schema.Follow{}.Mixin()
	followMixinFields0 := followMixin[0].Fields()
	_ = followMixinFields0
	followFields := schema.Follow{}.Fields()
	_ = followFields
	// followDescCreateTime is the schema descriptor for create_time field.
	followDescCreateTime := followMixinFields0[0].Descriptor()
	// follow.DefaultCreateTime holds the default value on creation for the create_time field.
	follow.DefaultCreateTime = followDescCreateTime.Default.(func() time.Time)
	tagMixin := schema.Tag{}.Mixin()
	tagMixinFields0 := tagMixin[0].Fields()
	_ = tagMixinFields0
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreateTime is the schema descriptor for create_time field.
	tagDescCreateTime := tagMixinFields0[0].Descriptor()
	// tag.DefaultCreateTime holds the default value on creation for the create_time field.
	tag.DefaultCreateTime = tagDescCreateTime.Default.(func() time.Time)
	// tagDescTag is the schema descriptor for tag field.
	tagDescTag := tagFields[0].Descriptor()
	// tag.TagValidator is a validator for the "tag" field. It is called by the builders before save.
	tag.TagValidator = tagDescTag.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescBio is the schema descriptor for bio field.
	userDescBio := userFields[3].Descriptor()
	// user.DefaultBio holds the default value on creation for the bio field.
	user.DefaultBio = userDescBio.Default.(string)
	// userDescImage is the schema descriptor for image field.
	userDescImage := userFields[4].Descriptor()
	// user.DefaultImage holds the default value on creation for the image field.
	user.DefaultImage = userDescImage.Default.(string)
}
