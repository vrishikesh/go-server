package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Golang Server Example")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	s := http.Server{
		Addr: ":9090",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("hit the handler func")
		}),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %+v", err)
	}
}
