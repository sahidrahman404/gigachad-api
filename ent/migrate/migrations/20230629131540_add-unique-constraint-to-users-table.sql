-- Create index "user_username" to table: "users"
CREATE UNIQUE INDEX `user_username` ON `users` (`username`);
-- Create index "user_email" to table: "users"
CREATE UNIQUE INDEX `user_email` ON `users` (`email`);
