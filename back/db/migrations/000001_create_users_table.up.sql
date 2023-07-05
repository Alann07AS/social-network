-- file: 000001_create_users_table.up.sql

-- Create the "user" table
CREATE TABLE IF NOT EXISTS user (
    ID INTEGER PRIMARY KEY UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    date_of_birth TEXT NOT NULL,
    avatar_image_path TEXT,
    nickname TEXT,
    about_me TEXT,
    is_public_profile TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (first_name, last_name)
);

-- Create the trigger "set_nickname_default"
CREATE TRIGGER IF NOT EXISTS set_nickname_default
AFTER INSERT ON user
BEGIN
    UPDATE user
    SET nickname = NEW.first_name
    WHERE ID = NEW.ID AND NEW.nickname IS NULL;
END;

-- Create the trigger "set_random_avatar_default"
CREATE TRIGGER IF NOT EXISTS set_random_avatar_default
AFTER INSERT ON user
BEGIN
    UPDATE user
    SET avatar_image_path = CASE RANDOM() % 3
        WHEN 0 THEN 'img/avatar/sunflower.jpeg'
        WHEN 1 THEN 'img/avatar/poit.jpeg'
        WHEN 2 THEN 'img/avatar/sun.jpeg'
        ELSE ''
        END
    WHERE ID = NEW.ID AND NEW.avatar_image_path IS NULL;
END;

-- Create the trigger "update_at_user"
CREATE TRIGGER IF NOT EXISTS update_at_user
AFTER INSERT ON user
BEGIN
    UPDATE user
    SET updated_at = CURRENT_TIMESTAMP;
    WHERE ID = NEW.ID IS NULL;
END;