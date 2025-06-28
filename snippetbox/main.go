package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a form for creating a new snippet..."))
}

// wildcardSegmentsExampleHandler demonstrates how to handle wildcard segments in a URL.
func wildcardSegmentsExampleHandler(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	itemId := r.PathValue("itemId")
	w.Write([]byte("Category: " + category + ", Item ID: " + itemId))
	// http://localhost:4000/wildcard/aaa/123
	// This will output: Category: aaa, Item ID: 123
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	// mux.HandleFunc("/", home)    // Handle all other requests to "/"
	mux.HandleFunc("/{$}", home) // Restrict this route to exact matches / only
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/wildcard/{category}/{itemId}", wildcardSegmentsExampleHandler)

	// // you can use http.HandleFunc() to register a handler function for the "/" URL pattern
	// // it will register a default servemux for you, so you don't need to create one
	// http.HandleFunc("/", home)
	// http.HandleFunc("/snippet/view", snippetView)
	// http.HandleFunc("/snippet/create", snippetCreate)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)

	// // if you pass nil as the second parameter, it will use the default servemux
	// err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
