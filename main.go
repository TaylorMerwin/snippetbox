package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox!"))
}

func main() {
	// Use the http.NewServeMux() function to init a new servemux
	// Then register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new server.
	// Two params are passed in:
	// The TCP network address to listen on, and the servemux we just created
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
