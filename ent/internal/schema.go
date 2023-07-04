// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/sahidrahman404/gigachad-api/ent/schema","Package":"github.com/sahidrahman404/gigachad-api/ent","Schemas":[{"name":"Equipment","config":{"Table":""},"edges":[{"name":"exercises","type":"Exercise","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"image","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}]},{"name":"Exercise","config":{"Table":""},"edges":[{"name":"routine_exercises","type":"RoutineExercise","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"workout_logs","type":"WorkoutLog","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"users","type":"User","field":"user_id","ref_name":"exercises","unique":true,"inverse":true},{"name":"equipments","type":"Equipment","field":"muscles_group_id","ref_name":"exercises","unique":true,"inverse":true},{"name":"muscles_groups","type":"MusclesGroup","field":"muscles_group_id","ref_name":"exercises","unique":true,"inverse":true},{"name":"exercise_types","type":"ExerciseType","field":"exercise_type_id","ref_name":"exercises","unique":true,"inverse":true}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"image","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"how_to","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"equipment_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"muscles_group_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"exercise_type_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}}]},{"name":"ExerciseType","config":{"Table":""},"edges":[{"name":"exercises","type":"Exercise","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"properties","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}]},{"name":"MusclesGroup","config":{"Table":""},"edges":[{"name":"exercises","type":"Exercise","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"image","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}]},{"name":"Routine","config":{"Table":""},"edges":[{"name":"routine_exercises","type":"RoutineExercise","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"users","type":"User","field":"user_id","ref_name":"routines","unique":true,"inverse":true}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}]},{"name":"RoutineExercise","config":{"Table":""},"edges":[{"name":"routines","type":"Routine","field":"routine_id","ref_name":"routine_exercises","unique":true,"inverse":true},{"name":"exercises","type":"Exercise","field":"exercise_id","ref_name":"routine_exercises","unique":true,"inverse":true},{"name":"users","type":"User","field":"user_id","ref_name":"routine_exercises","unique":true,"inverse":true}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"rest_timer","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"sets","type":{"Type":3,"Ident":"*schematype.Sets","PkgPath":"github.com/sahidrahman404/gigachad-api/ent/schema/schematype","PkgName":"schematype","Nillable":true,"RType":{"Name":"Sets","Ident":"schematype.Sets","Kind":22,"PkgPath":"github.com/sahidrahman404/gigachad-api/ent/schema/schematype","Methods":{}}},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"routine_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"exercise_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}}]},{"name":"Token","config":{"Table":""},"edges":[{"name":"users","type":"User","field":"user_id","ref_name":"tokens","unique":true,"inverse":true}],"fields":[{"name":"hash","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"expiry","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"scope","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}]},{"name":"User","config":{"Table":""},"edges":[{"name":"tokens","type":"Token","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"exercises","type":"Exercise","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"routines","type":"Routine","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"workouts","type":"Workout","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"workout_logs","type":"WorkoutLog","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"routine_exercises","type":"RoutineExercise","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"hashed_password","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"activated","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":2,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"version","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":1,"default_kind":2,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["username"]},{"unique":true,"fields":["email"]}]},{"name":"Workout","config":{"Table":""},"edges":[{"name":"users","type":"User","field":"user_id","ref_name":"workouts","unique":true,"inverse":true},{"name":"workout_logs","type":"WorkoutLog","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"volume","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"reps","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"time","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"sets","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"image","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":8,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0}}]},{"name":"WorkoutLog","config":{"Table":""},"edges":[{"name":"users","type":"User","field":"user_id","ref_name":"workout_logs","unique":true,"inverse":true},{"name":"exercises","type":"Exercise","field":"exercise_id","ref_name":"workout_logs","unique":true,"inverse":true},{"name":"workouts","type":"Workout","field":"workout_id","ref_name":"workout_logs","unique":true,"inverse":true}],"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"sets","type":{"Type":3,"Ident":"*schematype.Sets","PkgPath":"github.com/sahidrahman404/gigachad-api/ent/schema/schematype","PkgName":"schematype","Nillable":true,"RType":{"Name":"Sets","Ident":"schematype.Sets","Kind":22,"PkgPath":"github.com/sahidrahman404/gigachad-api/ent/schema/schematype","Methods":{}}},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"exercise_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"workout_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}}]}],"Features":["privacy","schema/snapshot"]}`
