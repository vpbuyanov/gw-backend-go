CREATE TABLE IF NOT EXISTS "forum_topic"
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(100) NOT NULL,
    creator_id INTEGER      NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES "user" (id)
);
