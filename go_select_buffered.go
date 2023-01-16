package main

import (
	"fmt"
	"time"
)

func init() {
	start = time.Now()
}

func Select_Buffered() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "Value 1"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan2 <- "Value 2"

	// In this case we can get response from either of the chan 1 or chan 2, only one case will get executed
	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
