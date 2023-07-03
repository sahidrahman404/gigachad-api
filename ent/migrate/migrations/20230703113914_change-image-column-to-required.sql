-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_equipment" table
CREATE TABLE `new_equipment` (`id` text NOT NULL, `name` text NOT NULL, `image` text NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "equipment" to new temporary table "new_equipment"
INSERT INTO `new_equipment` (`id`, `name`, `image`) SELECT `id`, `name`, `image` FROM `equipment`;
-- Drop "equipment" table after copying rows
DROP TABLE `equipment`;
-- Rename temporary table "new_equipment" to "equipment"
ALTER TABLE `new_equipment` RENAME TO `equipment`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
