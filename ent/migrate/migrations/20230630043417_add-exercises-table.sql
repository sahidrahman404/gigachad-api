-- Create "exercises" table
CREATE TABLE `exercises` (`id` text NOT NULL, `exercise_name` text NOT NULL, `how_to` text NOT NULL, `user_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `exercises_users_exercises` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
