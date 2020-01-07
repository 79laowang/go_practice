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

func area(r Rectangle) {
    fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r Rectangle) area() {
    fmt.Printf("Area method result: %d\n", (r.length * r.width))
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
    e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
    e.salary = newAge
}

type address struct {
    city string
    state string
}

func (a address) fullAddress() {
    fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
    firstName string
    lastName  string
    address
}

func perimeter(r *Rectangle) {
    fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

// method
func (r *Rectangle) perimeter() {
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

type myInt int

// non-struct receivers
func (a myInt) add(b myInt) myInt {
    return a + b
}

func main() {
    fmt.Println(" To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package.\n")
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

    fmt.Println("\nPointer Receivers vs Value Receivers:")
    e := Employee{
        name: "Makr Andrew",
        salary: 5000,
        currency: "$",
    }
    fmt.Printf("Employee name before change:%s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee salary before change: %d", e.salary)
    (&e).changeAge(5001)
    fmt.Printf("\nEmployee salary after change: %d", e.salary)
    e.changeAge(5002)
    fmt.Printf("\nEmployee salary after change: %d", e.salary)
    
    fmt.Println("\n\nMedhods of anonymous struct fields:")
    p11 := person {
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city: "Los Angeles",
            state: "ShanDong",
        },
    }    
    p11.fullAddress()  // accessing fullAddress method of address struct

    // When a function has a value argument, it will accept only a value argument.
    // When a method has a value receiver, it will accept both pointer and value receivers
    fmt.Println("\n\nValue receivers in methods vs Value arguments in functions")
    r1 := Rectangle {
        length: 10,
        width: 5,
    }
    area(r1)
    r1.area()
    p12 := &r1
     /*
       compilation error, cannot use p (type *rectangle) as type rectangle 
       in argument to area  
    */
    //area(p12)
    //for convenience will be interpreted by Go as (*p).area() since area has a value receiver. 
    p12.area() //calling value receiver with a pointer

    fmt.Println("\n\nPointer receivers in methods vs Pointer arguments in functions,")
    p13 := &r1 // pointer to r
    perimeter(p13)
    p13.perimeter()
    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)
    // will be interpreted by the language as (&r).perimeter() for convenience.
    r.perimeter()   //calling pointer receiver with a value

    fmt.Println("\n\nMethods with non-struct receivers,")
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
