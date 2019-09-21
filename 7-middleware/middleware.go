package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func handle(rw http.ResponseWriter, r *http.Request) {
	rid := r.Header.Get("request-id")
	if rid != "" {
		fmt.Fprintf(rw, "Request ID: %s\n", rid)
	}

	fmt.Fprintln(rw, "Hello, DDD East Anglia")
}

// Functions are, as mentioned in earlier examples, first class citizens in Go.
// This means that functions can be passed as values and returned as values. Allowing
// us to create middleware functions that can wrap a HTTP `HandlerFunc`. In this case
// we're adding a unique `request-id` to the HTTP headers of the incoming request,
// then passing the modified request into the handler function.
func addReqID(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		r.Header.Add("request-id", string(out))

		h(rw, r)
	}
}

func main() {
	http.Handle("/", addReqID(handle))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
