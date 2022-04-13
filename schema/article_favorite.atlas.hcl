table "article_favorites" {
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
  column "article_id" {
    null = false
    type = bigint
  }
  column "user_id" {
    null = false
    type = bigint
  }
  foreign_key "article_favorites_articles_fk" {
    columns     = [column.article_id]
    ref_columns = [table.articles.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
  foreign_key "article_favorites_users_fk" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
  primary_key {
    columns = [column.id]
  }
}
