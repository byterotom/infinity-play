-- tech enum
CREATE TYPE tech AS ENUM ('html', 'flash');

-- game table
CREATE TABLE game (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    technology tech NOT NULL,
    release_date DATE DEFAULT CURRENT_DATE NOT NULL,
    likes INTEGER DEFAULT 0 NOT NULL,
    votes INTEGER DEFAULT 0 NOT NULL
);

-- tags table
CREATE TABLE tags (
    tag_id SERIAL PRIMARY KEY,
    tag TEXT UNIQUE NOT NULL
);

-- game_tags table (relation table)
CREATE TABLE game_tags (
    game_id TEXT NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY (game_id, tag_id),
    FOREIGN KEY (game_id) REFERENCES game(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(tag_id) ON DELETE CASCADE
);

-- role_type enum
CREATE TYPE role_type AS ENUM ('master-admin', 'admin');

-- admin table
CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role role_type NOT NULL
);