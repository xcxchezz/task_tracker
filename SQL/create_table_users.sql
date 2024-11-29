CREATE TABLE IF NOT EXISTS public.users (
    id BIGINT NOT NULL PRIMARY KEY,
    username VARCHAR(32) NOT NULL,
    email VARCHAR(128) NOT NULL,
    password_hash VARCHAR(64) NOT NULL
);
