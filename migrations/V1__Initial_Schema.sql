CREATE DATABASE IF NOT EXISTS snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER IF NOT EXISTS 'web'@'%';
GRANT SELECT, INSERT, UPDATE, DELETE, EXECUTE ON snippetbox.* TO 'web'@'%';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'%' IDENTIFIED BY 'pass';