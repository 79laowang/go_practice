package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

//For Waitgroup
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Gorouting", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // The counter is decremented every call
}

type Job struct {
	id         int
	randomno   int
}

type Result struct {
	job          Job
	sumofdigits  int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)


func main() {
	fmt.Println("\nBuffered channels:")
	//Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.
	ch := make(chan string, 2)
	ch <- "Bob"
	ch <- "Wang"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch1 := make(chan int, 2)
	go write(ch1)
	time.Sleep(2 * time.Second)
	for v := range ch1 {
		fmt.Println("read value", v, "from ch1")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Dead lock:")
	ch2 := make(chan string, 2)
	ch2 <- "naveen"
	ch2 <- "paul"
	// ch2 <- "bob" // dead lock
	// write 3 strings to a buffered channel of capacity 2, the third write is blocked since the channel has exceeded its capacity
	// no concurrent routine reading from this channel. Hence there will be a deadlock and the program will panic at run time
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	fmt.Println("\nLength vs Capacity:")
	ch3 := make(chan string, 3)
	ch3 <- "naveen"
	ch3 <- "bob"
	fmt.Println("capacity is", cap(ch3))
	fmt.Println("length is", len(ch3))
	fmt.Println("read value", <-ch3)
	fmt.Println("new length is", len(ch3))
	
	fmt.Println("\nWaitGourp:")
	no := 3
	var wg sync.WaitGroup
	// the WaitGroup's counter is incremented by add method
	for i := 0; i < no; i++ {
		fmt.Println("Increase one go routine", i + 1)
		wg.Add(1)
		go process(i, &wg)
	}
	// This makes the main Goroutine to wait until the counter becomes zero
	wg.Wait()
	fmt.Println("All go routines finished execting!")

	/*
	In general, a worker pool is a collection of threads which are waiting for tasks to be assigned to them. 
	Once they finish the task assigned, they make themselves available again for the next task.
	*/
	fmt.Println("\nWorker Pool:")

}
