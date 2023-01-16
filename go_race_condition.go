package main

import (
	"fmt"
	"sync"
)

// goroutine increment global variable i
func worker(wg *sync.WaitGroup) {
	//  i = i + 1 calculation has 3 steps
	// (1) get value of i
	// (2) increment value of i by 1
	// (3) update value of i with new value
	// we have not put any lock here
	i = i + 1
	wg.Done()
}

func Go_race_demo() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	// wait until all 1000 gorutines are done
	wg.Wait()

	// value of i should be 1000
	fmt.Println("value of i after 1000 operations is", i)
}
