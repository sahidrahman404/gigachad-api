-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_routine_exercises" table
CREATE TABLE `new_routine_exercises` (`id` text NOT NULL, `rest_timer` integer NULL, `sets` json NOT NULL, `exercise_id` text NULL, `routine_id` text NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `routine_exercises_exercises_routine_exercises` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_routines_routine_exercises` FOREIGN KEY (`routine_id`) REFERENCES `routines` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_users_routine_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "routine_exercises" to new temporary table "new_routine_exercises"
INSERT INTO `new_routine_exercises` (`id`, `rest_timer`, `exercise_id`, `routine_id`, `user_id`) SELECT `id`, `rest_timer`, `exercise_id`, `routine_id`, `user_id` FROM `routine_exercises`;
-- Drop "routine_exercises" table after copying rows
DROP TABLE `routine_exercises`;
-- Rename temporary table "new_routine_exercises" to "routine_exercises"
ALTER TABLE `new_routine_exercises` RENAME TO `routine_exercises`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
