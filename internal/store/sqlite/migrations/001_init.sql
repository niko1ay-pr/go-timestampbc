-- +goose Up
CREATE TABLE IF NOT EXISTS polls (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status in('upcoming','active','ended')),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ballots (
    id TEXT PRIMARY KEY,
    poll_id TEXT NOT NULL,
    flat_id TEXT NOT NULL,
    random_num INTEGER NOT NULL,
    hash TEXT NOT NULL UNIQUE,
    answers TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status in ('pending', 'confirmed', 'invalid')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    confirmed_at TIMESTAMP NULL,
    block_hash TEXT NULL,
    transaction_id TEXT NULL,

    FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE,
    UNIQUE(poll_id, flat_id)
);

CREATE TABLE IF NOT EXISTS questions (
    id TEXT PRIMARY KEY,
    poll_id TEXT NOT NULL,
    title TEXT NOT NULL,
    type TEXT NOT NULL CHECK (type in ('single', 'multiple')),
    sequence INT NOT NULL DEFAULT 0,

    FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS question_options (
    id TEXT PRIMARY KEY,
    question_id TEXT NOT NULL,
    title TEXT NOT NULL,
    sequence INT NOT NULL DEFAULT 0,

    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);


CREATE INDEX IF NOT EXISTS idx_polls_status ON polls(status);
CREATE INDEX IF NOT EXISTS idx_questions_poll_id ON questions(poll_id);
CREATE INDEX IF NOT EXISTS idx_question_options_question_id ON question_options(question_id);
CREATE INDEX IF NOT EXISTS idx_ballot_poll_id ON ballots(poll_id);
CREATE INDEX IF NOT EXISTS idx_ballots_status ON ballots(status);
CREATE INDEX IF NOT EXISTS idx_ballots_flat_id ON ballots(flat_id);

-- +goose Down
DROP TABLE IF EXISTS polls;
DROP TABLE IF EXISTS ballots;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS question_options;
