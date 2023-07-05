-- file: 000001_create_users_table.down.sql

-- Drop the trigger "set_random_avatar_default"
DROP TRIGGER IF EXISTS set_random_avatar_default;

-- Drop the trigger "set_nickname_default"
DROP TRIGGER IF EXISTS set_nickname_default;

-- Drop the "user" table
DROP TABLE IF EXISTS user;

-- Drop the trigger "update_at_user"
DROP TRIGGER IF EXISTS update_at_user;

