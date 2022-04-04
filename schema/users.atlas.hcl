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
