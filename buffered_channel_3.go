package main

import (
	"fmt"
	"runtime"
)

func Squares_2(c chan int) {
	for i := 0; i < 6; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func Buffered_3() {
	fmt.Println("main() started")
	c := make(chan int, 3)
	go Squares_2(c)

	fmt.Println("active goroutines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())

	//go Squares_2(c)

	//fmt.Println("active goroutines", runtime.NumGoroutine())

	c <- 5
	c <- 6
	c <- 7
	c <- 8 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())
	fmt.Println("main() stopped")
}
