-- +goose Up

create table feeds (
    id uuid PRIMARY KEY,
    created timestamp not null,
    modified timestamp not null,
    name text not null,
    url text unique not null,
    user_id uuid not null REFERENCES  users(id) on delete cascade
);

-- +goose Down
drop table feeds;
