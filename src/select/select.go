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

func server3(ch chan string) {
	ch <- "from server3"
}

func server4(ch chan string) {
	ch <- "from server4"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process succesful"
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

    fmt.Println("\n--------Random select ---------\n")
	output3 := make(chan string)
	output4 := make(chan string)
	go server3(output3)
	go server4(output4)
	time.Sleep(1 * time.Second)
	select {
	case s3 := <-output3:
		fmt.Println(s3)
	case s4 := <-output4:
		fmt.Println(s4)
	}
	
	fmt.Println("\n---- default case in a select statement ----\n")
	ch1 := make(chan string)
	go process(ch1)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch1:
			fmt.Println("received value: ",v)
			return
		default:
			fmt.Println("no value received")
		}
	}


}
