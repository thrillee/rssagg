-- +goose Up

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created TIMESTAMP not null,
    modified TIMESTAMP not null,
    title text not null,
    description text null,
    published_at timestamp not null,
    url text not null unique,
    feed_id uuid not null references feeds(id) on delete cascade
);

-- +goose Down
DROP TABLE posts;
