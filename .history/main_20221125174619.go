package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Golang Server Example")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// s := http.Server{
	// 	Addr: ":9090",
	// 	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		log.Println("hit the handler func")
	// 	}),
	// }

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("server stopped: %+v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
