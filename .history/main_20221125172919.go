package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Golang Server Example")

	s := http.Server{
		Addr: ":9090",
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %+v", err)
	}
}
