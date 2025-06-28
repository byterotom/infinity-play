-- name: AddGame :one
INSERT INTO
    Game(id, name, description, technology, game_url)
VALUES
    (?, ?, ?, ?, ?) RETURNING *;

-- name: GetGameIdByName :one
SELECT
    id
FROM
    Game
WHERE
    name = ?;

-- name: DeleteGameById :one
DELETE FROM
    Game
WHERE
    id = ? RETURNING *;

-- name: GetGameByName :one
SELECT
    *
FROM
    Game
WHERE
    name = ?;

-- name: GetNewGames :many
SELECT
    *
FROM
    Game
ORDER BY
    release_date
LIMIT
    10;

-- name: GetTopRatedGames :many
SELECT
    *
FROM
    Game
ORDER BY
    (likes / votes) DESC
LIMIT
    10;

-- name: GetPopularGames :many
SELECT
    *
FROM
    Game
ORDER BY
    votes DESC
LIMIT
    10;