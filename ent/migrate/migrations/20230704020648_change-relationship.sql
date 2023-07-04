-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_routine_exercises" table
CREATE TABLE `new_routine_exercises` (`id` text NOT NULL, `rest_timer` integer NULL, `sets` json NOT NULL, `routine_id` text NOT NULL, `exercise_id` text NOT NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `routine_exercises_routines_routines` FOREIGN KEY (`routine_id`) REFERENCES `routines` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_exercises_exercises` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`) ON DELETE CASCADE, CONSTRAINT `routine_exercises_users_routine_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
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
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
