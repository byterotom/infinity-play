-- name: AddGame :exec
INSERT INTO
    game(id, name, description, technology)
VALUES
    ($1, $2, $3, $4);

-- name: AddNewTags :exec
INSERT INTO
    tags(tag)
VALUES
    ($1) ON CONFLICT DO NOTHING;

-- name: AddGameTags :exec
INSERT INTO
    game_tags(game_id, tag_id)
VALUES
    ($1, $2);

-- name: GetTagIdByName :one
SELECT
    tag_id
FROM
    tags
WHERE
    tag = $1;

-- name: GetGameIdByName :one
SELECT
    id
FROM
    game
WHERE
    name = $1;

-- name: DeleteGameById :exec
DELETE FROM
    game
WHERE
    id = $1;

-- name: GetGameByName :one
SELECT
    *
FROM
    game
WHERE
    name = $1;

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
    (likes :: float / NULLIF(votes, 0)) DESC NULLS LAST
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
    t.tag = $1;

-- name: GetGamesByPattern :many
SELECT
    *
FROM
    game
WHERE
    $1::text IS NOT NULL
    AND (
        name ILIKE '%' || $1::text || '%'
        OR description ILIKE '%' || $1::text || '%'
    );

-- name: GetAdminIdByCredentials :one
SELECT 
    id 
FROM 
    admin 
WHERE 
    username=$1 AND password=$2;

-- name: VoteGameById :exec
UPDATE game SET votes = votes+1 WHERE id = $1;

-- name: LikeGameById :exec
UPDATE game SET likes = likes+1 WHERE id = $1;