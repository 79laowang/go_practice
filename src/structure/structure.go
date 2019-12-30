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

    fmt.Println("\nAccessing individual fields of a struct:")
    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Println("Salary: $%d", emp6.salary)
   
    fmt.Println("\nCreate zero struct, assign values later:")
    var emp7 Employee
    emp7.firstName = "Jack"
    emp7.lastName = "Adams"
    fmt.Println("Employee 7:", emp7)

    fmt.Println("\nPointers to a struct:")
    emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", (*emp8).firstName)
    fmt.Println("Age:", (*emp8).age)
    fmt.Println("\nPointers to a struct another way:")
    fmt.Println("First Name:", emp8.firstName)
    fmt.Println("Age:", emp8.age)

    fmt.Println("\nAnonymous fields:")
    type Person struct {
        string
        int
    }
    p := Person{"Naveen", 50}
    fmt.Println("Person:", p)
    var p1 Person
    p1.string = "naveen"
    p1.int = 50
    fmt.Println("Struct with anonymous fileds:", p1)
}
