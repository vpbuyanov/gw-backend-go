CREATE TABLE IF NOT EXISTS "user" (
      id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
      name VARCHAR(50) NOT NULL,
      email VARCHAR(100) NOT NULL,
      hash_pass VARCHAR(1000) NOT NULL,
      is_admin BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "forum_topic" (
     id SERIAL PRIMARY KEY,
     title VARCHAR(100) NOT NULL,
     creator_id uuid NOT NULL,
     FOREIGN KEY(creator_id) REFERENCES "user"(id)
);

CREATE TABLE IF NOT EXISTS "comment" (
     id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
     creator_id uuid NOT NULL,
     topic_id INTEGER NOT NULL,
     text VARCHAR(500) NOT NULL,
     FOREIGN KEY(creator_id) REFERENCES "user"(id),
     FOREIGN KEY(topic_id) REFERENCES "forum_topic"(id)
);