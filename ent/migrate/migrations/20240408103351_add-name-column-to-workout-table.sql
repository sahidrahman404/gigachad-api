-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_workouts" table
CREATE TABLE `new_workouts` (`id` text NOT NULL, `volume` real NOT NULL, `duration` text NOT NULL, `sets` integer NOT NULL, `created_at` datetime NOT NULL, `image` json NULL, `description` text NULL, `name` text NOT NULL DEFAULT (''), `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `workouts_users_workouts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "workouts" to new temporary table "new_workouts"
INSERT INTO `new_workouts` (`id`, `volume`, `duration`, `sets`, `created_at`, `image`, `description`, `user_id`) SELECT `id`, `volume`, `duration`, `sets`, `created_at`, `image`, `description`, `user_id` FROM `workouts`;
-- Drop "workouts" table after copying rows
DROP TABLE `workouts`;
-- Rename temporary table "new_workouts" to "workouts"
ALTER TABLE `new_workouts` RENAME TO `workouts`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
