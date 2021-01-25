package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testStore struct {
	movieId int64
	movies  []*Movie
}

func (t testStore) Open() error {
	return nil
}

func (t testStore) Close() error {
	return nil
}

func (t testStore) GetMovies() ([]*Movie, error) {
	return t.movies, nil
}

func (t testStore) GetMovieById(id int64) (*Movie, error) {
	for _, m := range t.movies {
		if m.ID == id {
			return m, nil
		}
	}
	return nil, nil
}

func (t testStore) CreateMovie(m *Movie) error {
	t.movieId++
	m.ID = t.movieId
	t.movies = append(t.movies, m)
	return nil
}

func TestMovieCreateUnit(t *testing.T) {
	// Create test with Test DB
	srv := newServer()
	srv.store = &testStore{}

	// Prepare json body
	p := struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerUrl  string `json:"trailer_url"`
	}{
		Title:       "Inception",
		ReleaseDate: "2012-04-05",
		Duration:    145,
		TrailerUrl:  "https://www.youtube.com/watch?v=b-kTeJhHOhc",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	assert.Nil(t, err)

	r := httptest.NewRequest("POST", "/api/movies/", &buf)
	w := httptest.NewRecorder()

	srv.handleMovieCreate()(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMovieCreateIntegration(t *testing.T) {
	// Create test with Test DB
	srv := newServer()
	srv.store = &testStore{}

	// Prepare json body
	p := struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerUrl  string `json:"trailer_url"`
	}{
		Title:       "Inception",
		ReleaseDate: "2012-04-05",
		Duration:    145,
		TrailerUrl:  "https://www.youtube.com/watch?v=b-kTeJhHOhc",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	assert.Nil(t, err)

	r := httptest.NewRequest("POST", "/api/movies/", &buf)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTE1ODMyOTAsImlhdCI6MTYxMTU3OTY5MCwidXNlcm5hbWUiOiJnb2xhbmcifQ.t-Sh2Z2XWxnOSYQ5OKS2OmD0IbxU3JOBEnlGu_Qhx9M"
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	w := httptest.NewRecorder()

	srv.serveHttp(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
