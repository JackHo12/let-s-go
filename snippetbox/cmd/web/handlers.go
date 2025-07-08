package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// use Header.Add() to add a custom header to the response.
	w.Header().Add("Server", "Go Web Server")
	files := []string{
		"./ui/html/base.tmpl", // base template must be first
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error 1", http.StatusInternalServerError)
		return
	}
	// err = ts.Execute(w, nil)
	err = ts.ExecuteTemplate(w, "base", nil) // Execute the base template
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error 2", http.StatusInternalServerError)
	}
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

func downloadHandler(w http.ResponseController, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/file.zip")
}
