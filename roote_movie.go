package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type jsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerUrl  string `json:"trailer_url"`
}

func (s *server) handleMovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := s.store.GetMovies()
		if err != nil {
			log.Printf("Cannot load movies. err=%v\n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = make([]jsonMovie, len(movies))
		for i, m := range movies {
			resp[i] = mapMovieToJson(m)
		}

		s.respond(w, r, resp, http.StatusOK)
	}
}

func (s *server) handleMovieDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Cannot parse id to integer. err=%v", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}
		m, err := s.store.GetMovieById(id)
		if err != nil {
			log.Printf("Cannot load movie. err=%v", err)
			s.respond(w, r, m, http.StatusInternalServerError)
			return
		}
		var resp = mapMovieToJson(m)
		s.respond(w, r, resp, http.StatusOK)
	}
}

func (s *server) handleMovieCreate() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerUrl  string `json:"trailer_url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := s.decode(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse movie body err=%v", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}
		// Create a movie
		m := &Movie{
			ID:          0,
			Title:       req.Title,
			ReleaseDate: req.ReleaseDate,
			Duration:    req.Duration,
			TrailerUrl:  req.TrailerUrl,
		}
		// Store a movie in database
		err = s.store.CreateMovie(m)
		if err != nil {
			log.Printf("Cannot create movie in DB. err=%v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapMovieToJson(m)
		s.respond(w, r, resp, http.StatusOK)
	}
}

func mapMovieToJson(m *Movie) jsonMovie {
	return jsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerUrl:  m.TrailerUrl,
	}
}
