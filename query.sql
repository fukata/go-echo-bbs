-- name: GetThreadMessages :many
select * from thread_messages order by id desc limit $1;

-- name: CreateThreadMessage :one
insert into thread_messages (message, created_at) values ($1, $2) RETURNING *;
