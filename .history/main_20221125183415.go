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
		if r.Method == http.MethodGet {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "hello")
	})

	mux.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, fmt.Sprintf("Parse err %v", err), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "hello")
	})

	return mux
}
