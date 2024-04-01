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
table "equipment_exercises" {
  schema = schema.main
  column "equipment_id" {
    null = false
    type = text
  }
  column "exercise_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.equipment_id, column.exercise_id]
  }
  foreign_key "equipment_exercises_exercise_id" {
    columns     = [column.exercise_id]
    ref_columns = [table.exercises.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "equipment_exercises_equipment_id" {
    columns     = [column.equipment_id]
    ref_columns = [table.equipment.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "exercise_type_exercises" {
  schema = schema.main
  column "exercise_type_id" {
    null = false
    type = text
  }
  column "exercise_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.exercise_type_id, column.exercise_id]
  }
  foreign_key "exercise_type_exercises_exercise_id" {
    columns     = [column.exercise_id]
    ref_columns = [table.exercises.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "exercise_type_exercises_exercise_type_id" {
    columns     = [column.exercise_type_id]
    ref_columns = [table.exercise_types.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "muscles_group_exercises" {
  schema = schema.main
  column "muscles_group_id" {
    null = false
    type = text
  }
  column "exercise_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.muscles_group_id, column.exercise_id]
  }
  foreign_key "muscles_group_exercises_exercise_id" {
    columns     = [column.exercise_id]
    ref_columns = [table.exercises.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "muscles_group_exercises_muscles_group_id" {
    columns     = [column.muscles_group_id]
    ref_columns = [table.muscles_groups.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "tokens" {
  schema = schema.main
  column "id" {
    null = false
    type = text
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
    null = false
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
    null = true
    type = json
  }
  primary_key {
    columns = [column.id]
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
    type = json
  }
  primary_key {
    columns = [column.id]
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
  column "image" {
    null = true
    type = json
  }
  column "how_to" {
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
  column "user_id" {
    null = false
    type = text
  }
  column "workout_id" {
    null = false
    type = text
  }
  column "exercise_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "workout_logs_exercises_exercises" {
    columns     = [column.exercise_id]
    ref_columns = [table.exercises.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "workout_logs_workouts_workouts" {
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
  index "workoutlog_workout_id_exercise_id" {
    unique  = true
    columns = [column.workout_id, column.exercise_id]
  }
}
table "workouts" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "volume" {
    null = false
    type = integer
  }
  column "duration" {
    null = false
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
    null = true
    type = json
  }
  column "description" {
    null = true
    type = text
  }
  column "user_id" {
    null = false
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
table "routine_exercises" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "rest_time" {
    null = true
    type = text
  }
  column "sets" {
    null = false
    type = json
  }
  column "routine_id" {
    null = false
    type = text
  }
  column "exercise_id" {
    null = false
    type = text
  }
  column "user_id" {
    null = false
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
  foreign_key "routine_exercises_exercises_exercises" {
    columns     = [column.exercise_id]
    ref_columns = [table.exercises.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "routine_exercises_routines_routines" {
    columns     = [column.routine_id]
    ref_columns = [table.routines.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "routineexercise_routine_id" {
    columns = [column.routine_id]
  }
  index "routineexercise_routine_id_exercise_id" {
    unique  = true
    columns = [column.routine_id, column.exercise_id]
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
  column "reminder_id" {
    null = true
    type = text
  }
  column "user_id" {
    null = false
    type = text
  }
  column "reminders" {
    null = true
    type = json
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
schema "main" {
}
