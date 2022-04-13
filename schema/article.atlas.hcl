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
  index "slug_idx" {
    columns = [
      column.slug
    ]
    unique = true
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
