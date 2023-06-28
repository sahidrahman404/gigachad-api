table "atlas_schema_revisions" {
  schema = schema.main
  column "version" {
    null = false
    type = text
  }
  column "description" {
    null = false
    type = text
  }
  column "type" {
    null    = false
    type    = integer
    default = 2
  }
  column "applied" {
    null    = false
    type    = integer
    default = 0
  }
  column "total" {
    null    = false
    type    = integer
    default = 0
  }
  column "executed_at" {
    null = false
    type = datetime
  }
  column "execution_time" {
    null = false
    type = integer
  }
  column "error" {
    null = true
    type = text
  }
  column "error_stmt" {
    null = true
    type = text
  }
  column "hash" {
    null = false
    type = text
  }
  column "partial_hashes" {
    null = true
    type = json
  }
  column "operator_version" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.version]
  }
}
table "tokens" {
  schema = schema.main
  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }
  column "hash" {
    null = false
    type = blob
  }
  column "expiry" {
    null = false
    type = text
  }
  column "scope" {
    null = false
    type = text
  }
  column "user_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "tokens_users_tokens" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "users" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "email" {
    null = false
    type = text
  }
  column "username" {
    null = false
    type = text
  }
  column "hashed_password" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = text
  }
  column "activated" {
    null    = false
    type    = bool
    default = false
  }
  column "version" {
    null    = false
    type    = integer
    default = 1
  }
  primary_key {
    columns = [column.id]
  }
}
schema "main" {
}
