package main

import "fmt"

type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerUrl  string `db:"trailer_url"`
}

func (m Movie) String() string {
	return fmt.Sprintf("id=%v, title=%v, releaseDate=%v, duration=%v, trailerUrl=%v",
		m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerUrl)
}
