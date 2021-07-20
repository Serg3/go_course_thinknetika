// Package 'model' uses like a types storage across an application
package model

// Movie struct presents such type
type Movie struct {
	ID          int
	Name        string
	ReleaseYear int
	Rating      string
	Fee         int
	StudioID    int
}
