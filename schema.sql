-- We'll use UUIDs as primary keys for a variety of reasons, more atomic, etc.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- We'll need a table for users, with username and email as an indexable key
CREATE TABLE users (
  id   uuid PRIMARY KEY,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  image TEXT NOT NULL DEFAULT '',
  bio TEXT NOT NULL DEFAULT ''
);

CREATE INDEX idx_users_username_email
ON users(username, email);