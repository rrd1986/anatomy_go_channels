package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c)
}

func Unbuffered() {
	fmt.Println("Unbuffered started")

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
