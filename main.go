package main

import (
	"fmt"
)

func Greet(c chan string) {
	fmt.Println("Hello" + <-c + "!")
}

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c)
}

func main() {
	//Buffered()
	//Buffered_2()
	//Buffered_3()
	//Anonymous_routines()
	//Channel_of_channel()
	//Select_Usages()
	//Select_with_buffered_default()
	//Test_WaitGroup()
	//Worker_Pool_Demo()
	//Workers_with_waitgroup()
	//Go_race_demo()
	Mutex_avoiding_race_condition()

}
