package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct {
	router *mux.Router
	store  Store
}

const JWT_APP_KEY = "training.go"

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *server) serveHttp(w http.ResponseWriter, r *http.Request) {
	logRequestMiddleware(s.router.ServeHTTP).ServeHTTP(w, r)
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
