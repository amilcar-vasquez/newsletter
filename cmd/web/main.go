// Filename: cmd/web/main.go
// Write a sample Hello World programs
package main

import (
	"log"
	"net/http"
)

//handler function for home

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our newsletter"))
}

func main() {
	// mux is our router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
