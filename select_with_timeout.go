package main

import (
	"fmt"
	"time"
)

func init() {
	start = time.Now()
}

func service11(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service22(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func Select_with_timeout() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
