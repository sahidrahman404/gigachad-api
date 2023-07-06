-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_tokens" table
CREATE TABLE `new_tokens` (`id` text NOT NULL, `hash` blob NOT NULL, `expiry` text NOT NULL, `scope` text NOT NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `tokens_users_tokens` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "tokens" to new temporary table "new_tokens"
INSERT INTO `new_tokens` (`id`, `hash`, `expiry`, `scope`, `user_id`) SELECT `id`, `hash`, `expiry`, `scope`, `user_id` FROM `tokens`;
-- Drop "tokens" table after copying rows
DROP TABLE `tokens`;
-- Rename temporary table "new_tokens" to "tokens"
ALTER TABLE `new_tokens` RENAME TO `tokens`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
