-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS courses (
    id serial PRIMARY KEY,
    author_id int NOT NULL references users(id),
    course_image_id text,
    title varchar(250) not null,
    description text,
    start_date timestamp,
    end_date timestamp,
    category text[],
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    price float not null default 0,
    status varchar(15) not null default 'personal'
);

CREATE INDEX IF NOT EXISTS  idx_courses_title ON courses(title);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE courses;
DROP INDEX idx_courses_title;
-- +goose StatementEnd
