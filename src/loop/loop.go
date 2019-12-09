package main

import "fmt"

func main(){
    for i := 1; i <= 10; i++{
        if i > 7{
            break
        }
        if i % 2 ==0 {
            fmt.Printf(" %d", i)
        } else {
            continue
        }
    }
    fmt.Printf("\nline after for loop.\n")
}
