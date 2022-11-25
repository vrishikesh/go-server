package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Golang Server Example")

	s := http.Server{
		Addr:    ":9090",
		Handler: routes(),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %+v", err)
	}
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
