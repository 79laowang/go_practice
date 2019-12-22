package main

import (
    "fmt"
)

func main() {
    b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Printf("address of b is\n", a)

    a1 := 25
    var b1 *int
    if b1 == nil {
        fmt.Println("\nb1 is", b1)
        b1 = &a1
        fmt.Println("b1 after initialization is", b1)
    }

    fmt.Println("Creating pointers using the new function:")
    size := new(int)
    fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size,size)
    *size = 85
    fmt.Println("New size value is", *size)

    fmt.Println("Dereferencing a pointer:")
    b2 := 255
    a2 := &b2
    fmt.Println("Address of b2 is", a2)
    fmt.Println("value of b2 is", *a2)
    *a2++
    fmt.Println("New value of b2 is", b2)

    fmt.Println("\nPassing pointer to a function:")

}
