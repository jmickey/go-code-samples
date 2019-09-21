package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// Gorilla Mux is part of the Gorilla Toolkit (https://gorillatoolkit.org/)
	"github.com/gorilla/mux"
)

var defaultName = "DDD East Anglia"

// Gorilla Mux allows us to separate out GET and POST functions by using the .Methods
// method as can be seen on L#49 and L#50.
func greetingHandler(rw http.ResponseWriter, r *http.Request) {
	// We can retreive dynamic path variables using the `mux.Vars()` function provided
	// by Gorilla Mux. This allows us to template dynamic URLs safely, and provides
	// increased reability. See L#48 to see this in action.
	vars := mux.Vars(r)
	if _, ok := vars["name"]; !ok {
		fmt.Fprintf(rw, "Hello, %s", defaultName)
		return
	}
	fmt.Fprintf(rw, "Hello, %s", vars["name"])
}

func updateDefaultNameHandler(rw http.ResponseWriter, r *http.Request) {
	body := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, "error decoding body: %w", http.StatusBadRequest)
	}
	defer r.Body.Close()

	if name, ok := body["name"]; ok {
		defaultName = name
		rw.WriteHeader(http.StatusOK)
		return
	}

	http.Error(rw, "Invalid request body", http.StatusBadRequest)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", greetingHandler)
	router.HandleFunc("/hello", greetingHandler).Methods(http.MethodGet)
	router.HandleFunc("/hello", updateDefaultNameHandler).Methods(http.MethodPost)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
