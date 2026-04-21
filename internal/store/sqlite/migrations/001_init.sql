CREATE TABLE IF NOT EXISTS elections (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('upcoming', 'active', 'ended')),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS votes (
    id TEXT PRIMARY KEY,
    election_id TEXT NOT NULL,
    candidate_id TEXT NOT NULL,
    random_num TEXT NOT NULL,
    hash TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL CHECK (status IN ('pending', 'confirmed', 'invalid')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    confirmed_at TIMESTAMP NULL,
    block_hash TEXT NULL,
    transaction_id TEXT NULL,

    FOREIGN KEY (election_id) REFERENCES elections(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_votes_election_id ON votes(election_id);
CREATE INDEX IF NOT EXISTS idx_votes_hash ON votes(hash);
CREATE INDEX IF NOT EXISTS idx_votes_status ON votes(status);
CREATE INDEX IF NOT EXISTS idx_elections_status ON elections(status);
