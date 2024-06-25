-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS  users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL unique,
    password VARCHAR(50) NOT NULL,
    avatar_id VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_login ON users  (login);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP INDEX  IF EXISTS idx_users_login;
-- +goose StatementEnd
