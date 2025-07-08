package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	// // mux.HandleFunc("/", home)    // Handle all other requests to "/"
	// mux.HandleFunc("/{$}", home) // Restrict this route to exact matches / only
	// mux.HandleFunc("/snippet/view/{id}", snippetView)
	// mux.HandleFunc("/snippet/create", snippetCreate)
	// mux.HandleFunc("/wildcard/{category}/{itemId}", wildcardSegmentsExampleHandler)

	// Use the http.FileServer() function to serve static files from the "./ui/static/"
	// directory. The http.StripPrefix() function is used to remove the "/static"
	// prefix from the URL path before passing it to the file server.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", neuter(fileServer)))

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

func neuter(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
