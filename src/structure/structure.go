package main

import (
    "fmt"
    "structure/computer"
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

    fmt.Println("\nNest struct:")
    type Address struct {
        city, state string
    }
    type Person1 struct {
        name string
        age int
        address Address
    }
    var p2 Person1
    p2.name = "Ke"
    p2.age = 25 
    p2.address = Address {
        city: "Beijing",
        state: "Beijing",
    }
    fmt.Println("Name:", p2.name)
    fmt.Println("Age:", p2.age)
    fmt.Println("City:", p2.address.city)
    fmt.Println("Age:", p2.address.state)

    fmt.Println("\nPromoted fields,")
    type Person2 struct {
        name string
        age int
        Address
    }
    var p3 Person2
    p3.name = "Bob"
    p3.age = 12
    p3.Address = Address {
        city: "Tongzhou",
        state: "Beijing",
    }
    fmt.Println("Name:", p3.name)
    fmt.Println("Age:", p3.age)
    fmt.Println("City:", p3.city) //city is promoted field
    fmt.Println("State:", p3.state) //state is promoted field

    fmt.Println("\n- Exported Structs and Fields,")
    var spec computer.Spec
    spec.Maker = "apple"
    spec.Price = 50000
    fmt.Println("Spec:", spec)

    type name struct {
        firstName string
        lastName string
    }

    fmt.Println("\n- Structs are value types and are comparable if each of their fields are comparable,")
    name1 := name{"Steve", "Jobs"}
    name2 := name{"Steve", "Jobs"}
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }
    name3 := name{firstName:"Steve", lastName:"Jobs"}
    name4 := name{}
    name4.firstName = "Steve"
    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }

    fmt.Println("\n- Struct variables are not comparable if they contain fields which are not comparable,")
    type image struct {
        data map[int]int
    }
    image1 := image{data: map[int]int{
        0: 155,
    }}
    image2 := image{data: map[int]int{
        0: 155,
    }}
    fmt.Println("image1 and image2 cannot be compared!")
    fmt.Println("image1",image1)
    fmt.Println("image2",image2)
}

