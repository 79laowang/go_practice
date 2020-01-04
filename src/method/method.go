package main

import (
    "fmt"
    "math"
)

type Employee struct {
    name  string
    salary int
    currency string
}

type Rectangle struct {
    length int
    width int
}

type Circle struct {
    radius float64
}
/* 
displaySalary()  method has Employee as the receiver type
*/
func (e Employee) displaySalary(){
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func (r Rectangle) Area() int {
    return r.length * r.width
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func main() {
    emp1 := Employee {
        name:    "Sam Adolf",
        salary:  5000,
        currency: "$",
    }
    emp1.displaySalary() //Calling displaySalary() method of Employee type

    r := Rectangle{
        length: 10,
        width: 5,
    }
    fmt.Printf("\nArea of rectangle: %d\n", r.Area())
    c := Circle{
        radius: 12,
    }
    fmt.Printf("\nArea of circle: %f", c.Area())
}
