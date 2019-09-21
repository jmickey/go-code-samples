package main

import (
	"fmt"
	"log"
)

func getGreeting(name string) (string, error) {
	if name != "DDD East Anglia" {
		return "", fmt.Errorf("unexpected name '%s', should be 'DDD East Anglia'", name)
	}
	return fmt.Sprintf("Hello, %s", name), nil
}

func main() {
	var greeting string
	var err error

	greeting, err = getGreeting("DDD Edinborough")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}
