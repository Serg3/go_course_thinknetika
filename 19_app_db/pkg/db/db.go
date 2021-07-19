// Package 'db' contains a DB interface
package db

import (
	"context"
	"go_course_thinknetika/19-app-db/pkg/model"
)

// DB interface defines the database contract
type DB interface {
	InsertMovies(context.Context, []model.Movie) error
	DeleteMovie(context.Context, int) error
	UpdateMovie(context.Context, model.Movie) error
	SelectMovies(context.Context, int) ([]model.Movie, error)
}
