-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS playlists (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    image_id TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    author_id int NOT NULL references users(id),
    courses_ids int[]
);

CREATE INDEX IF NOT EXISTS idx_playlist_title ON playlists(title);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS playlists;

DROP INDEX IF EXISTS idx_playlist_title;
-- +goose StatementEnd
