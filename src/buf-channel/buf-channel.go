package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Bob"
	ch <- "Wang"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
