package main

import (
    "fmt"
    "math"
    "unsafe"
)

func main(){
    var width, height int = 100, 50
    fmt.Println("width is", width, "height is", height)
    var (
        name = "Bob"
        age = 29
        gender = "male"
    )
    fmt.Println("My name is：", name, "age:", age, "gender: ", gender)
    name1, age1 := "Edward", 40 // Short hand declaration, newly declared
    fmt.Println("Your name is:", name1, "age: ", age1)

    a, b := 128.01, 45.77
    c := math.Min(a, b)
    fmt.Println("Minimum:", c)
    fmt.Printf("type of a is %T, size of a is: %d\n", a, unsafe.Sizeof(a)) //type and size of a

    fmt.Println("Boolean:")
    a1 := true
    b1 := false
    c1 := a1 && b1
    fmt.Println("a1:", a1, "b1:", b1, "c1:", c1)

    fmt.Println("String usage:")
    first := "科"
    last := "王"
    myname := last + " " + first
    fmt.Println("My name is:", myname)

    fmt.Println("Type convertion:")
    sum := int(a) + age
    fmt.Printf("%d = %f + %d\n", sum, a, age)

    var defaultName = "Sam"
    type myString string
    var customName myString = "Sam"
    // customName = defaultName Not allowed
}
