DROP INDEX IF EXISTS movies_release_year_name_idx, movies_actors_uniq_idx, movies_producers_uniq_idx;
DROP TABLE IF EXISTS movies, actors, producers, studios, movies_actors, movies_producers;
DROP TYPE IF EXISTS rating;

CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18');

CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    second_name VARCHAR(100) NOT NULL,
    birthday DATE NOT NULL
);

CREATE TABLE producers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    second_name VARCHAR(100) NOT NULL,
    birthday DATE NOT NULL
);

CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    release_year INTEGER NOT NULL CHECK ( release_year >= 1800 ),
    rating RATING NOT NULL,
    fee BIGINT NOT NULL DEFAULT 0,
    studio_id INTEGER REFERENCES studios
);
CREATE UNIQUE INDEX movies_release_year_name_idx ON movies (release_year, name);

CREATE TABLE movies_actors (
    movie_id INTEGER REFERENCES movies,
    actor_id INTEGER REFERENCES actors
);
CREATE UNIQUE INDEX movies_actors_uniq_idx ON movies_actors (movie_id, actor_id);


CREATE TABLE movies_producers (
    movie_id INTEGER REFERENCES movies,
    producer_id INTEGER REFERENCES producers
);
CREATE UNIQUE INDEX movies_producers_uniq_idx ON movies_producers (movie_id, producer_id);
