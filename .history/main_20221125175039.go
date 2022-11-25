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
	http.HandleFunc("/hello", helloHandler)

	mux := &http.ServeMux{}
	mux.Handle("/", fileServer)

	s := http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %+v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
