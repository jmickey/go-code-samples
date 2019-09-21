package main

import "fmt"

// The type of the parameter and/or return value much be provided within the
// function definition. Parameters are in the form of `<name> <type>`. For return
// values, just the type is required.
func getGreeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	var greeting = getGreeting("DDD East Anglia")
	fmt.Println(greeting)
}
