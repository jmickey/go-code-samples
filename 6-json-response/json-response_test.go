package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPizzaHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/pizza", nil)
	rw := httptest.NewRecorder()

	handler := http.HandlerFunc(pizzaHandler)
	handler.ServeHTTP(rw, req)

	result := rw.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 response, received: %d", result.StatusCode)
	}

	var pizza *Pizza
	err := json.NewDecoder(result.Body).Decode(&pizza)
	if err != nil {
		t.Fatalf("Received error when decoding JSON response: %v", err)
	}
	defer result.Body.Close()

	if pizza.Base != "Thin" {
		t.Errorf("Expected pizza.Base to be 'Thin', received: %s", pizza.Base)
	}
}
