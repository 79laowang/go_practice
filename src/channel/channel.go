package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("\nHello world goroutine with channel")
	done <- true // write to channel
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

// Only send channel
func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chn1 chan int) {
	for i := 0; i < 10; i++ {
		chn1 <- i
	}
	close(chn1)
}

func digits(number int, dchn1 chan int) {
	for number != 0 {
		digit := number % 10
		dchn1 <- digit
		number /= 10
	}
	close(dchn1)
}

func calcSquares1(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes1(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	done := make(chan bool)
	go hello(done)
	<-done // read from channel
	fmt.Println("main function")

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	fmt.Println("\nOnly send channel and Unidirectional channels:")
	chn1 := make(chan int)
	go sendData(chn1)
	fmt.Println("read from channel:", <-chn1)

	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received from channel:", v, ok)
	}

	ch1 := make(chan int)
	go producer(ch1)
	for v1 := range ch1 {
		fmt.Println("Received from channel using range loop:", v1)
	}

	sqrch1 := make(chan int)
	cubech1 := make(chan int)
	go calcSquares1(number, sqrch1)
	go calcCubes1(number, cubech1)
	squares1, cubes1 := <-sqrch1, <-cubech1
	fmt.Println("Final output", squares1+cubes1)

}
