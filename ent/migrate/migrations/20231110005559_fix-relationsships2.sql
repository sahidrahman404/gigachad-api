-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_routines" table
CREATE TABLE `new_routines` (`id` text NOT NULL, `name` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `routines_users_routines` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "routines" to new temporary table "new_routines"
INSERT INTO `new_routines` (`id`, `name`, `user_id`) SELECT `id`, `name`, `user_id` FROM `routines`;
-- Drop "routines" table after copying rows
DROP TABLE `routines`;
-- Rename temporary table "new_routines" to "routines"
ALTER TABLE `new_routines` RENAME TO `routines`;
-- Create "new_workouts" table
CREATE TABLE `new_workouts` (`id` text NOT NULL, `name` text NOT NULL, `volume` integer NOT NULL, `reps` integer NOT NULL, `time` text NULL, `sets` integer NOT NULL, `created_at` text NOT NULL, `image` text NULL, `description` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `workouts_users_workouts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "workouts" to new temporary table "new_workouts"
INSERT INTO `new_workouts` (`id`, `name`, `volume`, `reps`, `time`, `sets`, `created_at`, `image`, `description`, `user_id`) SELECT `id`, `name`, `volume`, `reps`, `time`, `sets`, `created_at`, `image`, `description`, `user_id` FROM `workouts`;
-- Drop "workouts" table after copying rows
DROP TABLE `workouts`;
-- Rename temporary table "new_workouts" to "workouts"
ALTER TABLE `new_workouts` RENAME TO `workouts`;
-- Create "new_tokens" table
CREATE TABLE `new_tokens` (`id` text NOT NULL, `hash` blob NOT NULL, `expiry` text NOT NULL, `scope` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `tokens_users_tokens` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "tokens" to new temporary table "new_tokens"
INSERT INTO `new_tokens` (`id`, `hash`, `expiry`, `scope`, `user_id`) SELECT `id`, `hash`, `expiry`, `scope`, `user_id` FROM `tokens`;
-- Drop "tokens" table after copying rows
DROP TABLE `tokens`;
-- Rename temporary table "new_tokens" to "tokens"
ALTER TABLE `new_tokens` RENAME TO `tokens`;
-- Create "new_routine_exercises" table
CREATE TABLE `new_routine_exercises` (`id` text NOT NULL, `rest_timer` text NULL, `sets` json NOT NULL, `routine_id` text NOT NULL, `exercise_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `routine_exercises_routines_routines` FOREIGN KEY (`routine_id`) REFERENCES `routines` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_exercises_exercises` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_users_routine_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "routine_exercises" to new temporary table "new_routine_exercises"
INSERT INTO `new_routine_exercises` (`id`, `rest_timer`, `sets`, `routine_id`, `exercise_id`, `user_id`) SELECT `id`, `rest_timer`, `sets`, `routine_id`, `exercise_id`, `user_id` FROM `routine_exercises`;
-- Drop "routine_exercises" table after copying rows
DROP TABLE `routine_exercises`;
-- Rename temporary table "new_routine_exercises" to "routine_exercises"
ALTER TABLE `new_routine_exercises` RENAME TO `routine_exercises`;
-- Create index "routineexercise_routine_id" to table: "routine_exercises"
CREATE UNIQUE INDEX `routineexercise_routine_id` ON `routine_exercises` (`routine_id`);
-- Create index "routineexercise_routine_id_exercise_id" to table: "routine_exercises"
CREATE UNIQUE INDEX `routineexercise_routine_id_exercise_id` ON `routine_exercises` (`routine_id`, `exercise_id`);
-- Create "new_workout_logs" table
CREATE TABLE `new_workout_logs` (`id` text NOT NULL, `sets` json NOT NULL, `created_at` text NOT NULL, `exercise_workout_logs` text NULL, `user_id` text NOT NULL, `workout_workout_logs` text NULL, PRIMARY KEY (`id`), CONSTRAINT `workout_logs_exercises_workout_logs` FOREIGN KEY (`exercise_workout_logs`) REFERENCES `exercises` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_users_workout_logs` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_workouts_workout_logs` FOREIGN KEY (`workout_workout_logs`) REFERENCES `workouts` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "workout_logs" to new temporary table "new_workout_logs"
INSERT INTO `new_workout_logs` (`id`, `sets`, `created_at`, `exercise_workout_logs`, `user_id`, `workout_workout_logs`) SELECT `id`, `sets`, `created_at`, `exercise_workout_logs`, `user_id`, `workout_workout_logs` FROM `workout_logs`;
-- Drop "workout_logs" table after copying rows
DROP TABLE `workout_logs`;
-- Rename temporary table "new_workout_logs" to "workout_logs"
ALTER TABLE `new_workout_logs` RENAME TO `workout_logs`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
