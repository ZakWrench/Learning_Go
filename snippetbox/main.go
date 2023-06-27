package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't,
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return, the handle
	// would keep executing and also write the "hello" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Define a home handler function which write a byte slice containing
	//"Hello" as the response body.
	w.Write([]byte("Hello from Snippetbox.zak.net"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	//Use the http.NewServerMux() function to intialize a new servemux, then
	// Register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
