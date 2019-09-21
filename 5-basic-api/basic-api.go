package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello, DDD East Anglia")
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
