-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created TIMESTAMP not null,
    modified TIMESTAMP not null,
    name VARCHAR(75) not null
);

-- +goose Down
DROP TABLE users;
