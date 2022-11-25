package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Golang Server Example")

	fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	// http.HandleFunc("/hello", helloHandler)

	mux := &http.ServeMux{}
	mux.Handle("/", fileServer)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	s := http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %+v", err)
	}
}
