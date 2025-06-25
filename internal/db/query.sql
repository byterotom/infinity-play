-- name: GetAll :many
SELECT
    *
FROM
    Game;

-- name: AddGame :one
INSERT INTO
    Game(id, name, description, technology, game_url)
VALUES
    (?, ?, ?, ?, ?) RETURNING *;

-- name: GetIdByName :one
SELECT
    id
FROM
    Game
WHERE
    name = ?;

-- name: DeleteById :one
DELETE FROM
    Game
where
    id = ? RETURNING *;