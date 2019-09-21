package main

import "fmt"

func getGreeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	var greeting = getGreeting("DDD East Anglia")
	fmt.Println(greeting)
}
