-- While we could use a traditional migration tool, our use case for RealWorld is
-- simple enough that we won't reap the benefits of needing to evolve and maintain

-- We'll need a table for users, with username and email as an indexable key
CREATE TABLE users
(
    -- With MySQL, we'll store UUIDs as bytes, using the UUID Go type to map them into structs
    id         BINARY(16) PRIMARY KEY,
    username   VARCHAR(255)  NOT NULL,
    email      VARCHAR(255)  NOT NULL,
    password   VARCHAR(255)  NOT NULL,
    image      VARCHAR(4096) NOT NULL DEFAULT '',
    bio        VARCHAR(4096) NOT NULL DEFAULT '',
    created_at TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username_email
    ON users (username, email);
