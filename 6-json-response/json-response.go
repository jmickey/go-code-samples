package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var pizza *Pizza

type Pizza struct {
	Base     string   `json:"base"`
	Sauce    string   `json:"sauce"`
	Toppings []string `json:"toppings"`
}

func pizzaHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&pizza)
		if err != nil {
			http.Error(rw, "couldn't decode json body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		rw.WriteHeader(http.StatusOK)
	}

	if r.Method == "GET" {
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
