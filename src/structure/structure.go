package main

import (
    "fmt"
)

type Employee struct {
    firstName, lastName string
    age, salary         int
}

func main() {
    // creating structure using field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Wang",
    }
    
    // Creating structure without using field names
    emp2 := Employee{"Thomas", "Li", 29, 800}

    fmt.Println("Employee 1:", emp1)
    fmt.Println("Employee 2:", emp2)

    fmt.Println("Creating anonymous structures:")
    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }
    fmt.Println("Employee 3", emp3)

    fmt.Println("Zero value of structure:")
    var emp4 Employee // zero valued structure
    fmt.Println("Employee 4:", emp4)

    emp5 := Employee{
        firstName: "Jhon",
        lastName: "Paul",
    }
    fmt.Println("Specify part values of structure:")
    fmt.Println("Employee 5", emp5)
}
