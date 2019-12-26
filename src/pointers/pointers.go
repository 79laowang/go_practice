package main

import (
    "fmt"
)

func change(val *int) {
    *val = 55
}

func hello() *int {
    i := 5
    return &i
}

func modify(arr *[3]int) {
    (*arr)[0] = 90
}

func modify1(arr *[3]int) {
    arr[0] = 88
}

func modify2(sls []int) {
    sls[0] = 87
}

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
    a3 := 58
    fmt.Println("value of a3 before function call is:", a3)
    b3 := &a3
    change(b3)
    fmt.Println("value of a3 after function call is:", a3)

    fmt.Println("Returning pointer from a function:")
    d := hello()
    fmt.Println("Value of d", *d)

    fmt.Println("\nPass a pointer to an array as a argument to a function:")
    // Although this way of passing a pointer to an array as a argument to a function and making modification to it works, it is not the idiomatic way of achieving this in Go 
    a4 := [3]int{89, 90, 91}
    modify(&a4)
    fmt.Println(a4)
    modify1(&a4)
    fmt.Println(a4)

    fmt.Println("\nPass slice as a argument to a function:")
    modify2(a4[:])
    fmt.Println(a4)

}
