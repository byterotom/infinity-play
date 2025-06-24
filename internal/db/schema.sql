-- game table
CREATE TABLE game (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    technology TEXT NOT NULL,
    release_date DATE DEFAULT (DATE('now')) NOT NULL,
    likes INTEGER DEFAULT 0 NOT NULL,
    votes INTEGER DEFAULT 0 NOT NULL,
    thumbnail_url TEXT NOT NULL,
    gif_url TEXT NOT NULL,
    game_url TEXT NOT NULL
);

-- tags table
CREATE TABLE tags(
    tag_id INTEGER PRIMARY KEY AUTOINCREMENT,
    tag TEXT UNIQUE NOT NULL
);

-- game_tags table (relation table)
CREATE TABLE game_tags(
    game_id TEXT NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY(game_id, tag_id),
    FOREIGN KEY(game_id) REFERENCES game(id),
    FOREIGN KEY(tag_id) REFERENCES tags(tag_id)
);