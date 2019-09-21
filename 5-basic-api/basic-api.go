package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function that prints to the ResponseWriter. The handler satisfies the
// function signature required to be passed to http.HandleFunc, which is fed to the
// underlying http.DefaultServeMux object.
//
// It takes a ResponseWriter, and a pointer to the Request object.
// ResponseWriter is used to provide a response to the HTTP request.
// The Request object refers to the incoming Request to the HTTP server,
// which contains details of the request, including body, URL, query params, etc.
func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello, DDD East Anglia")
}

func main() {
	// Go treats functions as first class citizens. We pass the handler function as a
	// value to HandlerFunc.
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
