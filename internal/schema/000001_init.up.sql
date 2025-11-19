CREATE TABLE teams (
    team_name TEXT PRIMARY KEY
);

CREATE TABLE users (
    user_id   TEXT PRIMARY KEY,
    username  TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    team_name TEXT NOT NULL REFERENCES teams(team_name)
);
CREATE TABLE pull_requests (
    pull_request_id   TEXT PRIMARY KEY,
    pull_request_name TEXT NOT NULL,
    assigned_reviewers TEXT[],
    author_id         TEXT NOT NULL REFERENCES users(user_id),
    status            TEXT NOT NULL,
    created_at        TIMESTAMP WITH TIME ZONE,
    merged_at         TIMESTAMP WITH TIME ZONE
);