-- name: AddGame :exec
INSERT INTO
    game(id, name, description, technology, game_url)
VALUES
    (?, ?, ?, ?, ?);

-- name: AddNewTags :exec
INSERT
    OR IGNORE INTO tags(tag)
VALUES
    (?);

-- name: AddGameTags :exec
INSERT INTO
    game_tags(game_id, tag_id)
VALUES
    (?, ?);

-- name: GetTagIdByName :one
SELECT
    tag_id
FROM
    tags
WHERE
    tag = ?;

-- name: GetGameIdByName :one
SELECT
    id
FROM
    game
WHERE
    name = ?;

-- name: DeleteGameById :exec
DELETE FROM
    game
WHERE
    id = ?;

-- name: GetGameByName :one
SELECT
    *
FROM
    game
WHERE
    name = ?;

-- name: GetNewGames :many
SELECT
    *
FROM
    game
ORDER BY
    release_date
LIMIT
    10;

-- name: GetTopRatedGames :many
SELECT
    *
FROM
    game
ORDER BY
    (likes / votes) DESC
LIMIT
    10;

-- name: GetPopularGames :many
SELECT
    *
FROM
    game
ORDER BY
    votes DESC
LIMIT
    10;

-- name: GetGamesByTag :many
SELECT
    g.*
FROM
    game AS g
    JOIN game_tags AS gt ON g.id = gt.game_id
    JOIN tags AS t ON gt.tag_id = t.tag_id
WHERE
    t.tag = ?;

-- name: GetGamesByPattern :many
SELECT
    *
FROM
    game
WHERE
    sqlc.arg('pattern') IS NOT NULL
    AND (
        name LIKE '%' || sqlc.arg('pattern') || '%'
        OR description LIKE '%' || sqlc.arg('pattern') || '%'
    );