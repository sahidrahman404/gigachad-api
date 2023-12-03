-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_workouts" table
CREATE TABLE `new_workouts` (`id` text NOT NULL, `name` text NOT NULL, `volume` integer NOT NULL, `reps` integer NOT NULL, `time` text NULL, `sets` integer NOT NULL, `created_at` text NOT NULL, `image` json NULL, `description` text NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `workouts_users_workouts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "workouts" to new temporary table "new_workouts"
INSERT INTO `new_workouts` (`id`, `name`, `volume`, `reps`, `time`, `sets`, `created_at`, `image`, `description`, `user_id`) SELECT `id`, `name`, `volume`, `reps`, `time`, `sets`, `created_at`, `image`, `description`, `user_id` FROM `workouts`;
-- Drop "workouts" table after copying rows
DROP TABLE `workouts`;
-- Rename temporary table "new_workouts" to "workouts"
ALTER TABLE `new_workouts` RENAME TO `workouts`;
-- Create "new_workout_logs" table
CREATE TABLE `new_workout_logs` (`id` text NOT NULL, `sets` json NOT NULL, `created_at` text NOT NULL, `user_id` text NOT NULL, `workout_id` text NOT NULL, `exercise_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `workout_logs_users_workout_logs` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_workouts_workouts` FOREIGN KEY (`workout_id`) REFERENCES `workouts` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_exercises_exercises` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "workout_logs" to new temporary table "new_workout_logs"
INSERT INTO `new_workout_logs` (`id`, `sets`, `created_at`, `user_id`) SELECT `id`, `sets`, `created_at`, `user_id` FROM `workout_logs`;
-- Drop "workout_logs" table after copying rows
DROP TABLE `workout_logs`;
-- Rename temporary table "new_workout_logs" to "workout_logs"
ALTER TABLE `new_workout_logs` RENAME TO `workout_logs`;
-- Create index "workoutlog_workout_id_exercise_id" to table: "workout_logs"
CREATE UNIQUE INDEX `workoutlog_workout_id_exercise_id` ON `workout_logs` (`workout_id`, `exercise_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
