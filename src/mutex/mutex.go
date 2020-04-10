package main

import (
	"fmt"
	"sync"
	// "time"
)

var x = 0
var y = 1
func increment(wg *sync.WaitGroup) {
	x = x + 1
	y = x - y
	wg.Done()
}

var x1 = 0
var y1 = 1
func increment1(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x1 = x1 + 1
	y1 = x1 - y1
	m.Unlock()
	wg.Done()
}

var x2 = 0
var y2 = 1
func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x2 = x2 + 1
	y2 = x2 - y2
	<- ch
	wg.Done()
}

func main() {
	fmt.Println("work with race condition:")
	var w sync.WaitGroup
	for i :=0; i < 1000;  i++ {
		w.Add(1)
		go increment(&w)
	}
	// Main thread wait for child threads completion
	w.Wait()
	fmt.Println("Final value of y:", y)

	var m sync.Mutex
	fmt.Println("\nSolving the race condition using mutex:")
	for j :=0; j < 1000; j++ {
		w.Add(1)
		go increment1(&w, &m)
	}
	w.Wait()
	fmt.Println("Final value of y1:", y1)

    fmt.Println("\nSolving the race condition using channel:")
	ch := make(chan bool, 1)
	for j :=0; j < 1000; j++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()
	fmt.Println("Final value of y2:", y2)

}
