package main

import (
	"fmt"
	"sync"
	"time"
)

func service5(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() // decrement counter
}

func Test_WaitGroup() {
	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment counter
		go service5(&wg, i)
	}

	wg.Wait() // blocks here
	fmt.Println("main() stopped")
}
