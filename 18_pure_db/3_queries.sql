-- list of films by studio name
SELECT movies.name, st.name as company
FROM movies
    JOIN studios st on movies.studio_id = st.id;

-- list of films by actor
SELECT movies.name
FROM movies
    JOIN movies_actors ma on movies.id = ma.movie_id
    JOIN actors a on ma.actor_id = a.id
WHERE (a.first_name || ' ' || a.second_name) = 'Morgan Freeman';

-- count of films for producer
SELECT count(*)
FROM movies
    JOIN movies_producers mp on movies.id = mp.movie_id
    JOIN producers p on p.id = mp.producer_id
WHERE (p.first_name || ' ' || p.second_name) = 'Some Guy 2';

-- list of films by several producers
SELECT DISTINCT movies.name
FROM movies JOIN movies_producers mp on movies.id = mp.movie_id
WHERE mp.producer_id IN (
    SELECT producers.id
    FROM producers
    WHERE (producers.first_name || ' ' || producers.second_name) IN ('Some Guy 2', 'Some Guy 4')
)
ORDER BY movies.name;

-- count of films for actor
SELECT count(*)
FROM movies
         JOIN movies_actors ma on movies.id = ma.movie_id
         JOIN actors a on ma.actor_id = a.id
WHERE (a.first_name || ' ' || a.second_name) = 'Morgan Freeman';

-- produsers and actors with 2+ movies
SELECT role, full_name, movies_count
FROM (
      SELECT
         'Actor' as role,
         (first_name || ' ' || second_name) as full_name,
         count(*) as movies_count
      FROM
           movies_actors JOIN actors a on a.id = movies_actors.actor_id
      GROUP BY (first_name || ' ' || second_name)
      HAVING count(*) > 1
    ) as t
UNION ALL
(
    SELECT
        'Producer' as role,
        (first_name || ' ' || second_name) as full_name,
        count(*) as movies_count
    FROM
        movies_producers JOIN producers p on p.id = movies_producers.producer_id
    GROUP BY (first_name || ' ' || second_name)
    HAVING count(*) > 1
)
ORDER BY movies_count DESC;

-- count of films with 1000+ fee
SELECT count(*)
FROM movies
WHERE fee > 1000;

-- count of producers with 1000+ fee
SELECT count(DISTINCT producer_id)
FROM movies_producers
WHERE movie_id IN (
    SELECT id
    FROM movies
    WHERE fee > 1000
);

-- unique last names of actors
SELECT DISTINCT second_name
FROM actors
ORDER BY second_name;

-- count of films with same names
SELECT count(*)
FROM (
    SELECT name
    FROM movies
    GROUP BY name
    HAVING count(*) > 1
) AS t;
