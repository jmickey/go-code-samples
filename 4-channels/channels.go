// Channels allow for communication between Goroutines, as well as the main program
// (which itself is executing within a goroutine!)
package main

import "fmt"

func hello(msg chan string) {
	// We receive the value from the channel using the `<-` operator, which we then assign
	// to a variable. This operation will block until a value has been received.
	m := <-msg
	fmt.Printf("Received message via Channel: %s", m)
}

func main() {
	// We use the `make` builtin function to create the channel. Channels are strongly
	// typed, and that type must be provided.
	msg := make(chan string)

	go hello(msg)

	// We can send data to the channel via the `->` operator. By default channels can
	// receive a single value at a time, and will cause the program to block if a value
	// has already been sent to the channel and has not been received. Think of it as
	// a First-In-First-Out queue with a capacity of 1.
	//
	// We can make goroutines with larger capacities by passing a size when we call
	// `make`. E.g. `make(chan string, 5)` will create a channel that can hold up to 5
	// values.
	msg <- "hello"
}
