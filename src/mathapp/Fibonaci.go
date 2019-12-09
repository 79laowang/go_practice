package main

/*
* 1. go build mymath
* 2. go build mathapp
*/

import(
    "mymath"
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) <2 {
        fmt.Println("Input number:")
        return
    }

    num, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("input must be number", err)
        return
    }

    fmt.Println("Fibonacci:", num, mymath.Fibonacci(num))
}
