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
    type    = integer
    default = 0
  }
  column "version" {
    null    = false
    type    = integer
    default = 1
  }
  primary_key {
    columns = [column.id]
  }
  index "user_username" {
    unique  = true
    columns = [column.username]
  }
  index "user_email" {
    unique  = true
    columns = [column.email]
  }
}
table "exercises" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "how_to" {
    null = false
    type = text
  }
  column "equipment_id" {
    null = true
    type = text
  }
  column "muscles_group_id" {
    null = true
    type = text
  }
  column "exercise_type_id" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "exercises_users_exercises" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "exercises_muscles_groups_exercises" {
    columns     = [column.muscles_group_id]
    ref_columns = [table.muscles_groups.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "exercises_exercise_types_exercises" {
    columns     = [column.exercise_type_id]
    ref_columns = [table.exercise_types.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "exercises_equipment_exercises" {
    columns     = [column.muscles_group_id]
    ref_columns = [table.equipment.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "equipment" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "image" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
table "exercise_types" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "properties" {
    null = false
    type = json
  }
  column "description" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
table "muscles_groups" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "image" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
table "routines" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "name" {
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
  foreign_key "routines_users_routines" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "routine_exercises" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "rest_timer" {
    null = true
    type = integer
  }
  column "set" {
    null = false
    type = json
  }
  column "exercise_id" {
    null = true
    type = text
  }
  column "routine_id" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "routine_exercises_users_routine_exercises" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "routine_exercises_routines_routine_exercises" {
    columns     = [column.routine_id]
    ref_columns = [table.routines.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "routine_exercises_exercises_routine_exercises" {
    columns     = [column.exercise_id]
    ref_columns = [table.exercises.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "workouts" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "volume" {
    null = false
    type = integer
  }
  column "reps" {
    null = false
    type = integer
  }
  column "time" {
    null = true
    type = text
  }
  column "sets" {
    null = false
    type = integer
  }
  column "created_at" {
    null = false
    type = text
  }
  column "image" {
    null = false
    type = text
  }
  column "description" {
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
  foreign_key "workouts_users_workouts" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "workout_logs" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "sets" {
    null = false
    type = json
  }
  column "created_at" {
    null = false
    type = text
  }
  column "exercise_id" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = text
  }
  column "workout_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "workout_logs_workouts_workout_logs" {
    columns     = [column.workout_id]
    ref_columns = [table.workouts.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "workout_logs_users_workout_logs" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "workout_logs_exercises_workout_logs" {
    columns     = [column.exercise_id]
    ref_columns = [table.exercises.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
schema "main" {
}
