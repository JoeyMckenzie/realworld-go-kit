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
