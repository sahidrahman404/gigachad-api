// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EquipmentColumns holds the columns for the "equipment" table.
	EquipmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
	}
	// EquipmentTable holds the schema information for the "equipment" table.
	EquipmentTable = &schema.Table{
		Name:       "equipment",
		Columns:    EquipmentColumns,
		PrimaryKey: []*schema.Column{EquipmentColumns[0]},
	}
	// ExercisesColumns holds the columns for the "exercises" table.
	ExercisesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
		{Name: "how_to", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// ExercisesTable holds the schema information for the "exercises" table.
	ExercisesTable = &schema.Table{
		Name:       "exercises",
		Columns:    ExercisesColumns,
		PrimaryKey: []*schema.Column{ExercisesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exercises_users_exercises",
				Columns:    []*schema.Column{ExercisesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ExerciseTypesColumns holds the columns for the "exercise_types" table.
	ExerciseTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "properties", Type: field.TypeJSON},
		{Name: "description", Type: field.TypeString},
	}
	// ExerciseTypesTable holds the schema information for the "exercise_types" table.
	ExerciseTypesTable = &schema.Table{
		Name:       "exercise_types",
		Columns:    ExerciseTypesColumns,
		PrimaryKey: []*schema.Column{ExerciseTypesColumns[0]},
	}
	// MusclesGroupsColumns holds the columns for the "muscles_groups" table.
	MusclesGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
	}
	// MusclesGroupsTable holds the schema information for the "muscles_groups" table.
	MusclesGroupsTable = &schema.Table{
		Name:       "muscles_groups",
		Columns:    MusclesGroupsColumns,
		PrimaryKey: []*schema.Column{MusclesGroupsColumns[0]},
	}
	// RoutinesColumns holds the columns for the "routines" table.
	RoutinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// RoutinesTable holds the schema information for the "routines" table.
	RoutinesTable = &schema.Table{
		Name:       "routines",
		Columns:    RoutinesColumns,
		PrimaryKey: []*schema.Column{RoutinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routines_users_routines",
				Columns:    []*schema.Column{RoutinesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoutineExercisesColumns holds the columns for the "routine_exercises" table.
	RoutineExercisesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "rest_timer", Type: field.TypeString, Nullable: true},
		{Name: "sets", Type: field.TypeJSON},
		{Name: "routine_id", Type: field.TypeString},
		{Name: "exercise_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// RoutineExercisesTable holds the schema information for the "routine_exercises" table.
	RoutineExercisesTable = &schema.Table{
		Name:       "routine_exercises",
		Columns:    RoutineExercisesColumns,
		PrimaryKey: []*schema.Column{RoutineExercisesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routine_exercises_routines_routines",
				Columns:    []*schema.Column{RoutineExercisesColumns[3]},
				RefColumns: []*schema.Column{RoutinesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "routine_exercises_exercises_exercises",
				Columns:    []*schema.Column{RoutineExercisesColumns[4]},
				RefColumns: []*schema.Column{ExercisesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "routine_exercises_users_routine_exercises",
				Columns:    []*schema.Column{RoutineExercisesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "routineexercise_routine_id",
				Unique:  false,
				Columns: []*schema.Column{RoutineExercisesColumns[3]},
			},
			{
				Name:    "routineexercise_routine_id_exercise_id",
				Unique:  true,
				Columns: []*schema.Column{RoutineExercisesColumns[3], RoutineExercisesColumns[4]},
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "hash", Type: field.TypeBytes},
		{Name: "expiry", Type: field.TypeString},
		{Name: "scope", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tokens_users_tokens",
				Columns:    []*schema.Column{TokensColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "hashed_password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeString},
		{Name: "activated", Type: field.TypeInt, Default: 0},
		{Name: "version", Type: field.TypeInt, Default: 1},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[2]},
			},
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// WorkoutsColumns holds the columns for the "workouts" table.
	WorkoutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "volume", Type: field.TypeInt},
		{Name: "duration", Type: field.TypeString},
		{Name: "sets", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeString},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeString},
	}
	// WorkoutsTable holds the schema information for the "workouts" table.
	WorkoutsTable = &schema.Table{
		Name:       "workouts",
		Columns:    WorkoutsColumns,
		PrimaryKey: []*schema.Column{WorkoutsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workouts_users_workouts",
				Columns:    []*schema.Column{WorkoutsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WorkoutLogsColumns holds the columns for the "workout_logs" table.
	WorkoutLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "sets", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "workout_id", Type: field.TypeString},
		{Name: "exercise_id", Type: field.TypeString},
	}
	// WorkoutLogsTable holds the schema information for the "workout_logs" table.
	WorkoutLogsTable = &schema.Table{
		Name:       "workout_logs",
		Columns:    WorkoutLogsColumns,
		PrimaryKey: []*schema.Column{WorkoutLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workout_logs_users_workout_logs",
				Columns:    []*schema.Column{WorkoutLogsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workout_logs_workouts_workouts",
				Columns:    []*schema.Column{WorkoutLogsColumns[4]},
				RefColumns: []*schema.Column{WorkoutsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workout_logs_exercises_exercises",
				Columns:    []*schema.Column{WorkoutLogsColumns[5]},
				RefColumns: []*schema.Column{ExercisesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "workoutlog_workout_id_exercise_id",
				Unique:  true,
				Columns: []*schema.Column{WorkoutLogsColumns[4], WorkoutLogsColumns[5]},
			},
		},
	}
	// EquipmentExercisesColumns holds the columns for the "equipment_exercises" table.
	EquipmentExercisesColumns = []*schema.Column{
		{Name: "equipment_id", Type: field.TypeString},
		{Name: "exercise_id", Type: field.TypeString},
	}
	// EquipmentExercisesTable holds the schema information for the "equipment_exercises" table.
	EquipmentExercisesTable = &schema.Table{
		Name:       "equipment_exercises",
		Columns:    EquipmentExercisesColumns,
		PrimaryKey: []*schema.Column{EquipmentExercisesColumns[0], EquipmentExercisesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equipment_exercises_equipment_id",
				Columns:    []*schema.Column{EquipmentExercisesColumns[0]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "equipment_exercises_exercise_id",
				Columns:    []*schema.Column{EquipmentExercisesColumns[1]},
				RefColumns: []*schema.Column{ExercisesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ExerciseTypeExercisesColumns holds the columns for the "exercise_type_exercises" table.
	ExerciseTypeExercisesColumns = []*schema.Column{
		{Name: "exercise_type_id", Type: field.TypeString},
		{Name: "exercise_id", Type: field.TypeString},
	}
	// ExerciseTypeExercisesTable holds the schema information for the "exercise_type_exercises" table.
	ExerciseTypeExercisesTable = &schema.Table{
		Name:       "exercise_type_exercises",
		Columns:    ExerciseTypeExercisesColumns,
		PrimaryKey: []*schema.Column{ExerciseTypeExercisesColumns[0], ExerciseTypeExercisesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exercise_type_exercises_exercise_type_id",
				Columns:    []*schema.Column{ExerciseTypeExercisesColumns[0]},
				RefColumns: []*schema.Column{ExerciseTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "exercise_type_exercises_exercise_id",
				Columns:    []*schema.Column{ExerciseTypeExercisesColumns[1]},
				RefColumns: []*schema.Column{ExercisesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MusclesGroupExercisesColumns holds the columns for the "muscles_group_exercises" table.
	MusclesGroupExercisesColumns = []*schema.Column{
		{Name: "muscles_group_id", Type: field.TypeString},
		{Name: "exercise_id", Type: field.TypeString},
	}
	// MusclesGroupExercisesTable holds the schema information for the "muscles_group_exercises" table.
	MusclesGroupExercisesTable = &schema.Table{
		Name:       "muscles_group_exercises",
		Columns:    MusclesGroupExercisesColumns,
		PrimaryKey: []*schema.Column{MusclesGroupExercisesColumns[0], MusclesGroupExercisesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "muscles_group_exercises_muscles_group_id",
				Columns:    []*schema.Column{MusclesGroupExercisesColumns[0]},
				RefColumns: []*schema.Column{MusclesGroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "muscles_group_exercises_exercise_id",
				Columns:    []*schema.Column{MusclesGroupExercisesColumns[1]},
				RefColumns: []*schema.Column{ExercisesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EquipmentTable,
		ExercisesTable,
		ExerciseTypesTable,
		MusclesGroupsTable,
		RoutinesTable,
		RoutineExercisesTable,
		TokensTable,
		UsersTable,
		WorkoutsTable,
		WorkoutLogsTable,
		EquipmentExercisesTable,
		ExerciseTypeExercisesTable,
		MusclesGroupExercisesTable,
	}
)

func init() {
	ExercisesTable.ForeignKeys[0].RefTable = UsersTable
	RoutinesTable.ForeignKeys[0].RefTable = UsersTable
	RoutineExercisesTable.ForeignKeys[0].RefTable = RoutinesTable
	RoutineExercisesTable.ForeignKeys[1].RefTable = ExercisesTable
	RoutineExercisesTable.ForeignKeys[2].RefTable = UsersTable
	TokensTable.ForeignKeys[0].RefTable = UsersTable
	WorkoutsTable.ForeignKeys[0].RefTable = UsersTable
	WorkoutLogsTable.ForeignKeys[0].RefTable = UsersTable
	WorkoutLogsTable.ForeignKeys[1].RefTable = WorkoutsTable
	WorkoutLogsTable.ForeignKeys[2].RefTable = ExercisesTable
	EquipmentExercisesTable.ForeignKeys[0].RefTable = EquipmentTable
	EquipmentExercisesTable.ForeignKeys[1].RefTable = ExercisesTable
	EquipmentExercisesTable.Annotation = &entsql.Annotation{}
	ExerciseTypeExercisesTable.ForeignKeys[0].RefTable = ExerciseTypesTable
	ExerciseTypeExercisesTable.ForeignKeys[1].RefTable = ExercisesTable
	ExerciseTypeExercisesTable.Annotation = &entsql.Annotation{}
	MusclesGroupExercisesTable.ForeignKeys[0].RefTable = MusclesGroupsTable
	MusclesGroupExercisesTable.ForeignKeys[1].RefTable = ExercisesTable
	MusclesGroupExercisesTable.Annotation = &entsql.Annotation{}
}
