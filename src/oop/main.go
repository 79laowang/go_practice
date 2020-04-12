package main

import (
	"fmt"
	"oop/employee"
)

func main() {
	fmt.Println("Go oop!")
	e := employee.Employee {
		FirstName:"Sam",
		LastName:  "Tangna",
		TotalLeaves:  30,
		LeavesTaken:  20,
	}
	e.LeavesRemaining()
}
