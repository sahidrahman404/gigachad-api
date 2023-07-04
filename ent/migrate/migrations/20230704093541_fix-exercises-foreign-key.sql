-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_exercises" table
CREATE TABLE `new_exercises` (`id` text NOT NULL, `name` text NOT NULL, `image` text NULL, `how_to` text NULL, `equipment_id` text NULL, `exercise_type_id` text NULL, `muscles_group_id` text NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `exercises_equipment_exercises` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `exercises_exercise_types_exercises` FOREIGN KEY (`exercise_type_id`) REFERENCES `exercise_types` (`id`) ON DELETE CASCADE, CONSTRAINT `exercises_muscles_groups_exercises` FOREIGN KEY (`muscles_group_id`) REFERENCES `muscles_groups` (`id`) ON DELETE CASCADE, CONSTRAINT `exercises_users_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "exercises" to new temporary table "new_exercises"
INSERT INTO `new_exercises` (`id`, `name`, `image`, `how_to`, `equipment_id`, `exercise_type_id`, `muscles_group_id`, `user_id`) SELECT `id`, `name`, `image`, `how_to`, `equipment_id`, `exercise_type_id`, `muscles_group_id`, `user_id` FROM `exercises`;
-- Drop "exercises" table after copying rows
DROP TABLE `exercises`;
-- Rename temporary table "new_exercises" to "exercises"
ALTER TABLE `new_exercises` RENAME TO `exercises`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
