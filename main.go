package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("GoFlix")

	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}

func run() error {
	port := os.Getenv("PORT")

	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()
	if err != nil {
		return err
	}

	http.HandleFunc("/", srv.serveHttp)
	log.Printf("Serving HTTP on port 9000")
	err = http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	defer srv.store.Close()
	return nil
	/*movies, err := srv.store.GetMovies()
	if err != nil {
		return err
	}
	fmt.Printf("movies=%v\n", movies)*/
}
