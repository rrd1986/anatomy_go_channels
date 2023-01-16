package main

import "fmt"

func Anonymous_routines() {
	fmt.Println("main() started")
	c := make(chan string)

	// launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "John"
	fmt.Println("main() stopped")
}
