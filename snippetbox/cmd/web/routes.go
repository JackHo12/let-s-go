package main

import (
	"net/http"
)

func (app *application) routes(staticDir string) *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)                                                    // Handle GET requests to "/"
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)                               // Handle GET requests
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)                                // Handle GET requests
	mux.HandleFunc("GET /wildcard/{category}/{itemId}", app.wildcardSegmentsExampleHandler) // Handle GET requests
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)                           // Handle POST requests
	return mux
}
