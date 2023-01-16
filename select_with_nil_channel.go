package main

import "fmt"

func service(c chan string) {
	c <- "response"
}

func Select_with_nil_channel() {
	fmt.Println("main() started")

	var chan1 chan string

	go service(chan1)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res)
	default:
		fmt.Println("No response")
	}

	fmt.Println("main() stopped")
}
