// Filename: cmd/web/main.go
// Write a sample Hello World programs
package main

import (
	"flag"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
)

//handler function for home

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/home.page.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := map[string]string{
		"Title":    "Tuesday | Thursday",
		"Greeting": "Mawnin",
		"Message":  "It's a great day",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// about page handler
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the about page"))
}

// signup
func signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the signup page"))
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//create a new structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// mux is our router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/signup", signup)

	//use the logger. note the key:value pairs
	logger.Info("Starting server on", "addr", *addr)
	err := http.ListenAndServe(*addr, http.HandlerFunc(home))

	//use the logger to log any errors
	logger.Error(err.Error())
	os.Exit(1)
}
