package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// use Header.Add() to add a custom header to the response.
	w.Header().Add("Server", "Go Web Server")
	w.Write([]byte("Hello from Snippetbox"))
	// curl -i localhost:4000/
	// HTTP/1.1 200 OK
	// Server: Go Web Server
	// Date: Mon, 30 Jun 2025 08:41:52 GMT
	// Content-Length: 21
	// Content-Type: text/plain; charset=utf-8

	// Hello from Snippetbox%
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// msg := fmt.Sprintf("Display a specific snippet with ID %d", id)
	// w.Write([]byte(msg))
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Cache-Control", "aaa")      // This will add a Cache-Control header
	w.Header().Set("Cache-Control", "no-cache") // This will overwrite the previous Cache-Control header
	w.Header().Add("Cache-Control", "bbb")      // This will add multiple Cache-Control headers
	w.Header().Del("Cache-Control")             // This will remove the Cache-Control header
	w.Write([]byte(`{"name":"Alex"}`))
	// curl -i localhost:4000/snippet/create
	// HTTP/1.1 200 ok
	// Date: Mon, 30 Jun 2025 07:49:02 GMT
	// Content-Length: 43
	// Content-Type: aaa
	// Create a form for creating a new snippet...%
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated) // Set the status code to 201 Created
	// This function would handle the POST request for creating a new snippet.
	w.Write([]byte("Snippet created successfully!"))
	// curl -i -d "" localhost:4000/snippet/create
	// HTTP/1.1 201 Created
	// Date: Mon, 30 Jun 2025 07:44:41 GMT
	// Content-Length: 29
	// Content-Type: text/plain; charset=utf-8
	// Snippet created successfully!%
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
	// // mux.HandleFunc("/", home)    // Handle all other requests to "/"
	// mux.HandleFunc("/{$}", home) // Restrict this route to exact matches / only
	// mux.HandleFunc("/snippet/view/{id}", snippetView)
	// mux.HandleFunc("/snippet/create", snippetCreate)
	// mux.HandleFunc("/wildcard/{category}/{itemId}", wildcardSegmentsExampleHandler)

	// To restrict a route to a specific HTTP method,
	// you can prefix the route pattern with the necessary HTTP method when declaring it.
	mux.HandleFunc("GET /{$}", home)                                                    // Handle GET requests to "/"
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)                               // Handle GET requests
	mux.HandleFunc("GET /snippet/create", snippetCreate)                                // Handle GET requests
	mux.HandleFunc("GET /wildcard/{category}/{itemId}", wildcardSegmentsExampleHandler) // Handle GET requests
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)                           // Handle POST requests
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
