CREATE TABLE IF NOT EXISTS "comment"
(
    id         uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    creator_id uuid         NOT NULL,
    topic_id   INTEGER      NOT NULL,
    text       VARCHAR(500) NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES "user" (id),
    FOREIGN KEY (topic_id) REFERENCES "forum_topic" (id)
);
