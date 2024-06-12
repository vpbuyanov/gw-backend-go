CREATE TABLE IF NOT EXISTS "user"
(
    id         uuid    DEFAULT gen_random_uuid() PRIMARY KEY,
    name       VARCHAR(50)  NOT NULL,
    email      VARCHAR(100) NOT NULL UNIQUE,
    phone      VARCHAR(11)  NOT NULL UNIQUE,
    hash_pass  VARCHAR(100) NOT NULL,
    is_admin   BOOLEAN DEFAULT FALSE,
    is_blocked BOOLEAN DEFAULT FALSE
);
