-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_routines" table
CREATE TABLE `new_routines` (`id` text NOT NULL, `name` text NOT NULL, `reminder_id` text NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `routines_users_routines` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "routines" to new temporary table "new_routines"
INSERT INTO `new_routines` (`id`, `name`, `user_id`) SELECT `id`, `name`, `user_id` FROM `routines`;
-- Drop "routines" table after copying rows
DROP TABLE `routines`;
-- Rename temporary table "new_routines" to "routines"
ALTER TABLE `new_routines` RENAME TO `routines`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
