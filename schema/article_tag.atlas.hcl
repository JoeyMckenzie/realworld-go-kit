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
