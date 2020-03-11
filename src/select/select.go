package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	fmt.Println("select usage:")
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	// The select statement blocks until one of its cases is ready, and then will terminate.
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
        fmt.Println(s2)
	}
}
