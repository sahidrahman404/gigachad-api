-- Create "tokens" table
CREATE TABLE `tokens` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `hash` blob NOT NULL, `expiry` text NOT NULL, `scope` text NOT NULL, `user_id` text NULL, CONSTRAINT `tokens_users_tokens` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `email` text NOT NULL, `username` text NOT NULL, `hashed_password` text NOT NULL, `name` text NOT NULL, `created_at` text NOT NULL, `activated` bool NOT NULL DEFAULT false, `version` integer NOT NULL DEFAULT 1, PRIMARY KEY (`id`));
