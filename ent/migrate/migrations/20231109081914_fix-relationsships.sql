-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_workout_logs" table
CREATE TABLE `new_workout_logs` (`id` text NOT NULL, `sets` json NOT NULL, `created_at` text NOT NULL, `exercise_workout_logs` text NULL, `user_id` text NULL, `workout_workout_logs` text NULL, PRIMARY KEY (`id`), CONSTRAINT `workout_logs_exercises_workout_logs` FOREIGN KEY (`exercise_workout_logs`) REFERENCES `exercises` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_users_workout_logs` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `workout_logs_workouts_workout_logs` FOREIGN KEY (`workout_workout_logs`) REFERENCES `workouts` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "workout_logs" to new temporary table "new_workout_logs"
INSERT INTO `new_workout_logs` (`id`, `sets`, `created_at`, `user_id`) SELECT `id`, `sets`, `created_at`, `user_id` FROM `workout_logs`;
-- Drop "workout_logs" table after copying rows
DROP TABLE `workout_logs`;
-- Rename temporary table "new_workout_logs" to "workout_logs"
ALTER TABLE `new_workout_logs` RENAME TO `workout_logs`;
-- Create "new_exercises" table
CREATE TABLE `new_exercises` (`id` text NOT NULL, `name` text NOT NULL, `image` text NULL, `how_to` text NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `exercises_users_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "exercises" to new temporary table "new_exercises"
INSERT INTO `new_exercises` (`id`, `name`, `image`, `how_to`, `user_id`) SELECT `id`, `name`, `image`, `how_to`, `user_id` FROM `exercises`;
-- Drop "exercises" table after copying rows
DROP TABLE `exercises`;
-- Rename temporary table "new_exercises" to "exercises"
ALTER TABLE `new_exercises` RENAME TO `exercises`;
-- Create "equipment_exercises" table
CREATE TABLE `equipment_exercises` (`equipment_id` text NOT NULL, `exercise_id` text NOT NULL, PRIMARY KEY (`equipment_id`, `exercise_id`), CONSTRAINT `equipment_exercises_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_exercises_exercise_id` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE);
-- Create "exercise_type_exercises" table
CREATE TABLE `exercise_type_exercises` (`exercise_type_id` text NOT NULL, `exercise_id` text NOT NULL, PRIMARY KEY (`exercise_type_id`, `exercise_id`), CONSTRAINT `exercise_type_exercises_exercise_type_id` FOREIGN KEY (`exercise_type_id`) REFERENCES `exercise_types` (`id`) ON DELETE CASCADE, CONSTRAINT `exercise_type_exercises_exercise_id` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE);
-- Create "muscles_group_exercises" table
CREATE TABLE `muscles_group_exercises` (`muscles_group_id` text NOT NULL, `exercise_id` text NOT NULL, PRIMARY KEY (`muscles_group_id`, `exercise_id`), CONSTRAINT `muscles_group_exercises_muscles_group_id` FOREIGN KEY (`muscles_group_id`) REFERENCES `muscles_groups` (`id`) ON DELETE CASCADE, CONSTRAINT `muscles_group_exercises_exercise_id` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
