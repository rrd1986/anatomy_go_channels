package main

import (
	"fmt"
	"time"
)

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func Select_Usages() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	// select does not work like loop when any case get a response it will get executed
	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		// Here before response from service 2 as it waits for 5 second the main thread will exit
		fmt.Println("Response from service 2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
