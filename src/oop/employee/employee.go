package employee

import (
	"fmt"
)

// Made the starting letter e of Employee struct to lower case, unexported the employee struct and prevented access from other packages.
type employee struct {
	FirstName    string
	LastName     string
	TotalLeaves  int
	LeavesTaken  int
}

func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeaves, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
