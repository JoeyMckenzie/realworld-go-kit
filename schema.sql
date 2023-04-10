-- We'll use UUIDs as primary keys for a variety of reasons, more atomic, etc.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- We'll need a table for users, with username and email as an indexable key
CREATE TABLE users
(
    id         uuid PRIMARY KEY   DEFAULT uuid_generate_v4(),
    username   TEXT      NOT NULL,
    email      TEXT      NOT NULL,
    password   TEXT      NOT NULL,
    image      TEXT      NOT NULL DEFAULT '',
    bio        TEXT      NOT NULL DEFAULT '',
    create_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username_email
    ON users (username, email);