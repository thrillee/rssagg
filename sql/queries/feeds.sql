-- name: CreateFeed :one
INSERT INTO feeds (
    id, created, modified, name, url, user_id
) VALUES ( $1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
select * from feeds order by created desc;


-- name: CreateFeedFollow :one
INSERT INTO feed_follows (
    id, created, modified, user_id, feed_id
) VALUES ( $1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserFeedFollows :many
select * from feed_follows where user_id=$1;

-- name: DeleteFeedFollow :exec
delete from feed_follows where id=$1 and user_id=$2;
