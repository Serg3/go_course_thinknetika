// package 'pg' is a movie management service
package pg

import (
	"context"
	"go_course_thinknetika/19-app-db/pkg/model"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PG is a type that implements functions for working with the movies DB
type PG struct {
	pool *pgxpool.Pool
}

// New creates a new PG object
func New(pool *pgxpool.Pool) *PG {
	p := PG{
		pool: pool,
	}
	return &p
}

// InsertMovies adds a new movie to DB via transaction
func (p *PG) InsertMovies(ctx context.Context, movies []model.Movie) error {
	sql := "INSERT INTO movies (name, release_year, rating, fee, studio_id) VALUES ($1, $2, $3, $4, $5)"
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	batch := new(pgx.Batch)
	for _, movie := range movies {
		batch.Queue(sql, movie.Name, movie.ReleaseYear, movie.Rating, movie.Fee, movie.StudioID)
	}

	results := tx.SendBatch(ctx, batch)
	err = results.Close()
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// DeleteMovie removes a movie from DB
func (p *PG) DeleteMovie(ctx context.Context, id int) error {
	_, err := p.pool.Exec(ctx, "DELETE FROM movies WHERE id = $1", id)
	return err
}

// UpdateMovie makes changes in the existing movie record
func (p *PG) UpdateMovie(ctx context.Context, movie model.Movie) error {
	sql := "UPDATE movies SET name = $1, release_year = $2, rating = $3, fee = $4, studio_id = $5 WHERE id = $6"
	_, err := p.pool.Exec(ctx, sql, movie.Name, movie.ReleaseYear, movie.Rating, movie.Fee, movie.StudioID, movie.ID)
	return err
}

// SelectMovies returns a list of movies by StudioID or all movies if StudioID is not provided (= 0)
func (p *PG) SelectMovies(ctx context.Context, StudioID int) ([]model.Movie, error) {
	var movies []model.Movie
	var err error
	var rows pgx.Rows

	rows, err = p.pool.Query(ctx, "SELECT * from MOVIES WHERE studio_id = $1 OR $1 = 0", StudioID)

	if err != nil {
		return movies, err
	}

	for rows.Next() {
		var m model.Movie
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.ReleaseYear,
			&m.Rating,
			&m.Fee,
			&m.StudioID,
		)
		if err != nil {
			return movies, err
		}
		movies = append(movies, m)
	}

	err = rows.Err()
	if err != nil {
		return movies, err
	}

	return movies, nil
}
