// Basic example of accepting dynamic URL paths using the http stdlib package.
// This isn't the nicest way to handle this usecase, and as such the next example
// will demonstate a better way to do this using a 3rd party library called Gorilla Mux.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	// Retreive the name from the path. This method of enabling dynamic paths is
	// pretty unreliable and prone to errors (e.g. what if the caller adds additional
	// paths? Then we'll end up with a pretty ugly response.)
	name := r.URL.Path[len("/view/"):]
	fmt.Fprintf(rw, "Hello, %s", name)
}

func main() {
	http.HandleFunc("/hello/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
