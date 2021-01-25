package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/token/", s.handleTokenCreate()).Methods("POST")
	s.router.HandleFunc("/api/movies/{id:[0-9]+}", s.loggedOnly(s.handleMovieDetail())).Methods("GET")
	s.router.HandleFunc("/api/movies/", s.loggedOnly(s.handleMovieList())).Methods("GET")
	s.router.HandleFunc("/api/movies/", s.loggedOnly(s.handleMovieCreate())).Methods("POST")
}
