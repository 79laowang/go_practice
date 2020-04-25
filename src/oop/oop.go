package main

import (
	"fmt"
	"oop/employee"
)

func main() {
	fmt.Println("Go oop!")
	e1 := employee.New("Bob", "Wang", 25, 15)
	e1.LeavesRemaining()
}
