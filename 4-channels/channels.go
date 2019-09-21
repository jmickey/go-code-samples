package main

import "fmt"

func hello(msg chan string) {
	m := <-msg
	fmt.Println(m)
}

func main() {
	msg := make(chan string)
	go hello(msg)
	msg <- "hello"
}
