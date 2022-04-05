schema "public" {
}

table "articles" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start     = 0
      increment = 1
    }
  }
  column "created_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "updated_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "title" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "slug" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "description" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "body" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "user_id" {
    null = false
    type = bigint
  }
  column "tags" {
    null = false
    type = sql("integer[]")
  }
  foreign_key "articles_users_fk" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
  primary_key {
    columns = [column.id]
  }
}

table "article_tags" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start     = 0
      increment = 1
    }
  }
  column "created_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "updated_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "tag_id" {
    null = false
    type = bigint
  }
  column "article_id" {
    null = false
    type = bigint
  }
  foreign_key "article_tags_articles_fk" {
    columns     = [column.article_id]
    ref_columns = [table.articles.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
  foreign_key "article_tags_tags_fk" {
    columns     = [column.tag_id]
    ref_columns = [table.tags.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
  primary_key {
    columns = [column.id]
  }
}

table "tags" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start     = 0
      increment = 1
    }
  }
  column "created_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "tag" {
    null    = false
    type    = character_varying
    default = ""
  }
  index "tag_idx" {
    columns = [
      column.tag
    ]
    unique = true
  }
  primary_key {
    columns = [column.id]
  }
}

table "user_profile_follows" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start     = 0
      increment = 1
    }
  }
  column "created_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "updated_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "follower_user_id" {
    null = false
    type = integer
  }
  column "followee_user_id" {
    null = false
    type = integer
  }
  primary_key {
    columns = [column.id]
  }
}

table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start     = 0
      increment = 1
    }
  }
  column "created_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "updated_at" {
    null = false
    type = timestamp_with_time_zone
  }
  column "username" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "email" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "password" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "bio" {
    null    = false
    type    = character_varying
    default = ""
  }
  column "image" {
    null    = false
    type    = character_varying
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
}

