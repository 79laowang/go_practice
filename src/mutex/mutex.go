package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	time.Sleep(250 * time.Millisecond)
	wg.Done()
}

func main() {
	fmt.Println("work with race condition:")
	var w sync.WaitGroup
	for i :=0; i < 10000;  i++ {
		w.Add(1)
		go increment(&w)
	}
	// Main thread wait for child threads completion
	w.Wait()
	fmt.Println("Final value of x:", x)
}
