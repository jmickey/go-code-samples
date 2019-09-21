package main

import (
	"fmt"
	"log"
)

// Go supports multiple return values, which must be wrapped in `()`.
//
// Errors are first-class values in Go. When an error occurs it is idiomatic to
// return the error back up the stack to the function caller, so they can choose
// how they wish to act on that error. In the case that an error does NOT occur,
// then retrun nil.
func getGreeting(name string) (string, error) {
	if name != "DDD East Anglia" {
		return "", fmt.Errorf("unexpected name '%s', should be 'DDD East Anglia'", name)
	}
	return fmt.Sprintf("Hello, %s", name), nil
}

func main() {
	// Variables can be declared and instantiated separately in the form of:
	// `var <name> <type>`.
	var greeting string
	var err error

	greeting, err = getGreeting("DDD Edinborough")
	// Check if an error was returned. If err != nil, then act on that error.
	// In this case we simply log the error and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}
