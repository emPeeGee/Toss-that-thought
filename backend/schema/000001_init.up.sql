CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- DROP TABLE users;
-- DROP TABLE thoughts;
-- DROP TABLE users_thoughts;

CREATE TABLE users
(
    id            serial       NOT NULL UNIQUE,
    email         varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    first_name    varchar(255) NOT NULL,
    last_name     varchar(255) NOT NULL,
    registered_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE thoughts (
                          id serial PRIMARY KEY unique ,
                          passphrase VARCHAR NOT NULL,
                          lifetime TIMESTAMP NOT NULL,
                          created_date TIMESTAMP DEFAULT current_timestamp,
                          is_burned BOOL DEFAULT FALSE,
                          time_accessed TIMESTAMP DEFAULT NULL,
                          metadata_key UUID default uuid_generate_v4 (),
                          thought_key uuid default uuid_generate_v4 ()
);
-- Excluded status and recepients

CREATE TABLE users_thoughts (
                                id serial unique,
                                user_id int references users(id) on delete cascade  not null,
                                thought_id int references thoughts (id) on delete cascade not null
);

INSERT INTO thoughts (passphrase, lifetime) VALUES
    ('It is a secret', timestamp '2022-11-11 15:00:00')
