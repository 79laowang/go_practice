package main

import (
    "fmt"
)
/*
Since a string is a slice of bytes, it's possible to access each byte of a string.
*/

func printBytes(s string) {
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])  //output Unicode, %x is the format specifier for hexadecimal.
    }
}

func printChars(s string) {
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c ", s[i])
    }
}

func main() {
    name := "Hello World"
    fmt.Println("Output with Unicode type:")
    printBytes(name)
    fmt.Println("Output with char type:")
    printChars(name)
}
