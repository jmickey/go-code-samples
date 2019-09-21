package main

import "fmt"

func task(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go task("World")
	task("Hello")
}
