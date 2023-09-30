-- name: GetUserByName :one
select * from users where name = $1 fetch first 1 row only;
