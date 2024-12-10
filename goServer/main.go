package main

import (
	"fmt"   // Provides formatted I/O functions.
	"log"   // Provides logging capabilities for errors or information.
	"net/http" // Implements HTTP client and server.

	// This program creates a simple HTTP server that serves static files,
	// handles a form submission, and responds to a "hello" endpoint.
)

// formHandler processes HTTP POST requests made to the "/form" endpoint.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the HTTP request.
	// ParseForm populates r.Form and r.PostForm from the request body.
	if err := r.ParseForm(); err != nil {
		// If there's an error during parsing, respond with an error message.
		fmt.Fprintf(w, "parseform err:%v", err)
		return
	}

	// Indicate to the user that the form submission was successful.
	fmt.Fprintf(w, "post was successful\n")

	// Extract the "name" and "address" values from the submitted form data.
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Respond with the extracted form values.
	fmt.Fprintf(w, "name is: %s\n", name)
	fmt.Fprintf(w, "address is: %s\n", address)
}

// helloHandler processes HTTP GET requests made to the "/hello" endpoint.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is exactly "/hello".
	if r.URL.Path != "/hello" {
		// Respond with a 404 error if the path doesn't match.
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Ensure the HTTP method is GET, as "/hello" only supports GET requests.
	if r.Method != "GET" {
		// Respond with an error if the method is not GET.
		http.Error(w, "method is not supported!", http.StatusMethodNotAllowed)
		return
	}

	// Respond with a simple "hello" message.
	fmt.Fprintf(w, "hello")
}

// main is the entry point of the program.
func main() {
	// Print a message to indicate the server is starting.
	fmt.Println("Starting server on port 8000...")

	// Serve static files from the "static" directory.
	// The FileServer handler serves files from the given filesystem root.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer) // Map the root URL path "/" to the file server.

	// Map the "/form" path to the formHandler function.
	http.HandleFunc("/form", formHandler)

	// Map the "/hello" path to the helloHandler function.
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on port 8000.
	// ListenAndServe listens on the specified address and calls the handlers for incoming requests.
	if err := http.ListenAndServe(":8000", nil); err != nil {
		// Log a fatal error if the server fails to start.
		log.Fatal(err)
	}
}





/*
	to run this we need to write : localhost:8000/hello localhost:8000/form.html
*/