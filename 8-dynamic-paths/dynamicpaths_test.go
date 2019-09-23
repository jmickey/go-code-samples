package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDynamicPaths(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello/Test", nil)
	rw := httptest.NewRecorder()

	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rw, req)

	result := rw.Result()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("Couldn't read response body: %v", err)
	}
	defer result.Body.Close()

	if string(body) != "Hello, Test" {
		t.Errorf("Invalid response body. Expected: 'Hello, Test', recieved: %s", string(body))
	}
}
