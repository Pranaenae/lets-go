CREATE USER 'web'@'%';
GRANT SELECT, INSERT, UPDATE, DELETE, EXECUTE ON snippetbox.* TO 'web'@'%';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'%' IDENTIFIED BY 'pass';