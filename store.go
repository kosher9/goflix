package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*Movie, error)
	GetMovieById(id int64) (*Movie, error)
	CreateMovie(m *Movie) error
}

type dbStore struct {
	db *sqlx.DB
}

var schema = `CREATE TABLE IF NOT EXISTS movie 
(
 id int(64) NOT NULL AUTO_INCREMENT,
 title varchar(256) NOT NULL,
 release_date varchar(256) NOT NULL,
 duration int(64) NOT NULL,
 trailer_url varchar(256) NOT NULL,
 PRIMARY KEY (id)
)`

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("mysql", "root:root@/goflix")
	if err != nil {
		return err
	}
	log.Println("Connected to db")
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}

func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (store *dbStore) GetMovieById(id int64) (*Movie, error) {
	var movie = &Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (store *dbStore) CreateMovie(m *Movie) error {
	res, err := store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?, ?, ?, ?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerUrl)
	if err != nil {
		return err
	}
	m.ID, err = res.LastInsertId()
	return err
}
