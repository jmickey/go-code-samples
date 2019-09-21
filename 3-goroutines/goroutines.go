// Goroutines are the core concurrency primitive in Go. They are effectively
// lightweight thread that execute concurrently to the code that launches them.
package main

import "fmt"

// Goroutines cannot return or receive input from stdin, but can still write
// to stdout/stderr.
func task(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	// Goroutines are launched via the `go` keyword. Which will launch the Goroutine
	// and continue to execute the program without any blocking.
	go task("World")
	task("Hello")
}
