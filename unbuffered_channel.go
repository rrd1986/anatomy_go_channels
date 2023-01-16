package main

import "fmt"

func Unbuffered() {
	//fmt.Printf("type of the channel and value of the channel is: %T , %v", c, c)
	fmt.Println("Unbuffered started")
	//c := make(chan string)

	//go greet(c)

	//c <- "Rashmi"

	c := make(chan int)
	go squares(c)

	for {
		//time.Sleep(1 * time.Second)
		val, ok := <-c
		if !ok {
			fmt.Println(val)
			break
		} else {
			fmt.Println(val, ok)
		}
	}

	// range operator can also be used instead of for/while loop
	s := make(chan int)
	go squares(s)

	for val := range s {
		fmt.Println(val)
	}

	fmt.Println("Unbuffered stopped")

}
