-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_users" table
CREATE TABLE `new_users` (`id` text NOT NULL, `email` text NOT NULL, `username` text NOT NULL, `hashed_password` text NOT NULL, `name` text NOT NULL, `unit` text NOT NULL DEFAULT ('METRIC'), `created_at` datetime NOT NULL, `activated` integer NOT NULL DEFAULT (0), `version` integer NOT NULL DEFAULT (1), PRIMARY KEY (`id`));
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `email`, `username`, `hashed_password`, `name`, `created_at`, `activated`, `version`) SELECT `id`, `email`, `username`, `hashed_password`, `name`, `created_at`, `activated`, `version` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "user_username" to table: "users"
CREATE UNIQUE INDEX `user_username` ON `users` (`username`);
-- Create index "user_email" to table: "users"
CREATE UNIQUE INDEX `user_email` ON `users` (`email`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
