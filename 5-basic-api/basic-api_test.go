package main

import (
	"net/http"
	// The httptest pkg provides various functions for testing HTTP handlers and servers.
	"net/http/httptest"
	"testing"
)

func TestBasicAPI(t *testing.T) {
	// Create a new Request object to test our handler.
	req := httptest.NewRequest("GET", "/", nil)

	// NewRecorder will allow us to record the reponse received from the handler.
	// It satisfies the ResponseWriter interface.
	rw := httptest.NewRecorder()

	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rw, req)

	// Retreive the result from the ResponseWriter.
	result := rw.Result()

	if result.StatusCode != http.StatusOK {
		// t.Error will mark the test as failed.
		t.Errorf("Expected status code to be 200, received: %d", result.StatusCode)
	}
}
