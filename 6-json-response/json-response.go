package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var pizza *Pizza

// Pizza is a struct that defines a Pizza object. The json tags allow Go to map the
// JSON request body to the struct fields, allowing them to be populated when we decode
// the body.
type Pizza struct {
	Base     string   `json:"base"`
	Sauce    string   `json:"sauce"`
	Toppings []string `json:"toppings"`
}

func pizzaHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// We decode the JSON response be providing the `Decode` method with a pointer to a
		// nil struct. An error is returned if the decode fails.
		err := json.NewDecoder(r.Body).Decode(&pizza)
		if err != nil {
			http.Error(rw, "couldn't decode json body", http.StatusBadRequest)
			return
		}
		// The `defer` keyword in Go, as the name suggests, allows us to defer the closing
		// of the request body when the body is no longer within scope (when the function
		// returns). `defer` will also be called if the program exits or panics.
		defer r.Body.Close()

		rw.WriteHeader(http.StatusOK)
	}

	if r.Method == "GET" {
		// If no pizza was previously POST'd then return a margherita
		if pizza == nil {
			pizza = &Pizza{
				Base:  "Thin",
				Sauce: "Tomato",
				Toppings: []string{
					"Mozzarella",
					"Cherry Tomatos",
				},
			}
		}

		// We `Marshal` (encode) the struct into a JSON []byte object for transport over HTTP.
		resp, err := json.Marshal(&pizza)
		if err != nil {
			http.Error(rw, "couldn't encode response", http.StatusInternalServerError)
			return
		}

		_, err = rw.Write(resp)
		if err != nil {
			http.Error(rw, "failed to write response", http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/pizza", pizzaHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
