-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_exercises" table
CREATE TABLE `new_exercises` (`id` text NOT NULL, `name` text NOT NULL, `how_to` text NOT NULL, `equipment_id` text NULL, `muscles_group_id` text NULL, `exercise_type_id` text NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `exercises_equipment_exercises` FOREIGN KEY (`muscles_group_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `exercises_exercise_types_exercises` FOREIGN KEY (`exercise_type_id`) REFERENCES `exercise_types` (`id`) ON DELETE CASCADE, CONSTRAINT `exercises_muscles_groups_exercises` FOREIGN KEY (`muscles_group_id`) REFERENCES `muscles_groups` (`id`) ON DELETE CASCADE, CONSTRAINT `exercises_users_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "exercises" to new temporary table "new_exercises"
INSERT INTO `new_exercises` (`id`, `how_to`, `user_id`) SELECT `id`, `how_to`, `user_id` FROM `exercises`;
-- Drop "exercises" table after copying rows
DROP TABLE `exercises`;
-- Rename temporary table "new_exercises" to "exercises"
ALTER TABLE `new_exercises` RENAME TO `exercises`;
-- Create "equipment" table
CREATE TABLE `equipment` (`id` text NOT NULL, `name` text NOT NULL, `image` text NULL, PRIMARY KEY (`id`));
-- Create "exercise_types" table
CREATE TABLE `exercise_types` (`id` text NOT NULL, `name` text NOT NULL, `properties` json NOT NULL, `description` text NOT NULL, PRIMARY KEY (`id`));
-- Create "muscles_groups" table
CREATE TABLE `muscles_groups` (`id` text NOT NULL, `name` text NOT NULL, `image` text NOT NULL, PRIMARY KEY (`id`));
-- Create "routines" table
CREATE TABLE `routines` (`id` text NOT NULL, `name` text NOT NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `routines_users_routines` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create "routine_exercises" table
CREATE TABLE `routine_exercises` (`id` text NOT NULL, `rest_timer` integer NULL, `set` json NOT NULL, `exercise_id` text NULL, `routine_id` text NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `routine_exercises_exercises_routine_exercises` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_routines_routine_exercises` FOREIGN KEY (`routine_id`) REFERENCES `routines` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_users_routine_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create "workouts" table
CREATE TABLE `workouts` (`id` text NOT NULL, `name` text NOT NULL, `volume` integer NOT NULL, `reps` integer NOT NULL, `time` text NULL, `sets` integer NOT NULL, `created_at` text NOT NULL, `image` text NOT NULL, `description` text NOT NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `workouts_users_workouts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create "workout_logs" table
CREATE TABLE `workout_logs` (`id` text NOT NULL, `sets` json NOT NULL, `created_at` text NOT NULL, `exercise_id` text NULL, `user_id` text NULL, `workout_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `workout_logs_exercises_workout_logs` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_users_workout_logs` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_workouts_workout_logs` FOREIGN KEY (`workout_id`) REFERENCES `workouts` (`id`) ON DELETE CASCADE);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
