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
  column "updated_at" {
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
