// Filename: cmd/web/main.go
// Write a sample Hello World programs
package main

import (
	"flag"
	"log"
	"net/http"
)

//handler function for home

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our newsletter"))
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// mux is our router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, http.HandlerFunc(home))
	log.Fatal(err)
}
