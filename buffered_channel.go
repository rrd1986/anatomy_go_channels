package main

import (
	"fmt"
)

func Squares(c chan int) {
	for i := 0; i <= 4; i++ {
		num := <-c
		fmt.Println("Entered Squares")
		fmt.Println(num * num)
	}
}

func Buffered() {

	fmt.Println("Buffered() started")
	c := make(chan int, 3)

	go Squares(c)

	c <- 1
	fmt.Println("entered Buffered")
	c <- 2
	fmt.Println("entered Buffered")
	c <- 3
	fmt.Println("entered Buffered")
	// but since the buffer is not overflowing (as we didn’t push any new value), the main goroutine will not block and program exists.
	c <- 4
	fmt.Println("entered Buffered")
	c <- 5
	fmt.Println("entered Buffered")

	fmt.Println("Buffered() stopped")

}
