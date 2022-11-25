package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

func run() error {
	log.Println("Golang Server Example")

	s := http.Server{
		Addr:    ":9090",
		Handler: routes(),
	}

	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("server stopped: %+v", err)
	}

	return nil
}

func routes() *http.ServeMux {
	fileServer := http.FileServer(http.Dir("./static"))

	mux := &http.ServeMux{}
	mux.Handle("/", fileServer)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	return mux
}
