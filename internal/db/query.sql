-- name: GetAll :many
SELECT * FROM Game;

-- name: AddOne :one
INSERT INTO Game(
    id,name,description,technology,thumbnail_url,gif_url,game_url
) VALUES(
    ?,?,?,?,?,?,?
)
RETURNING *;