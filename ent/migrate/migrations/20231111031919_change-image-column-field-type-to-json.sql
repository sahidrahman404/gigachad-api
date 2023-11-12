-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_muscles_groups" table
CREATE TABLE `new_muscles_groups` (`id` text NOT NULL, `name` text NOT NULL, `image` json NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "muscles_groups" to new temporary table "new_muscles_groups"
INSERT INTO `new_muscles_groups` (`id`, `name`, `image`) SELECT `id`, `name`, `image` FROM `muscles_groups`;
-- Drop "muscles_groups" table after copying rows
DROP TABLE `muscles_groups`;
-- Rename temporary table "new_muscles_groups" to "muscles_groups"
ALTER TABLE `new_muscles_groups` RENAME TO `muscles_groups`;
-- Create "new_equipment" table
CREATE TABLE `new_equipment` (`id` text NOT NULL, `name` text NOT NULL, `image` json NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "equipment" to new temporary table "new_equipment"
INSERT INTO `new_equipment` (`id`, `name`, `image`) SELECT `id`, `name`, `image` FROM `equipment`;
-- Drop "equipment" table after copying rows
DROP TABLE `equipment`;
-- Rename temporary table "new_equipment" to "equipment"
ALTER TABLE `new_equipment` RENAME TO `equipment`;
-- Create "new_exercises" table
CREATE TABLE `new_exercises` (`id` text NOT NULL, `name` text NOT NULL, `image` json NOT NULL, `how_to` text NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `exercises_users_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "exercises" to new temporary table "new_exercises"
INSERT INTO `new_exercises` (`id`, `name`, `image`, `how_to`, `user_id`) SELECT `id`, `name`, `image`, `how_to`, `user_id` FROM `exercises`;
-- Drop "exercises" table after copying rows
DROP TABLE `exercises`;
-- Rename temporary table "new_exercises" to "exercises"
ALTER TABLE `new_exercises` RENAME TO `exercises`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
