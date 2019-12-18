package main

import (
    "fmt"
    "unicode/utf8"
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
    // support unicode character
    runes := []rune(s)
    for i := 0; i < len(runes); i++ {
        fmt.Printf("%c ", runes[i])
    }
}

func printCharsAndBytes(s string) {
    for index, rune := range s {
        fmt.Printf("%c starts at byte %d\n", rune, index)
    }
}

// Strings are immutable in Go. Once a string is created it's not possible to change it.  
// func mutate(s string)string {
    // any valid unicode character within single quote is a rune 
    // s[0] = 'a'
    // return s
// }

// To workaround this string immutability, strings are converted to a slice of runes. 
func mutate_slice(s []rune) string {
    s[0] = 'a'
    return string(s)
}

func length(s string) {
    fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func main() {
    name := "Hello World"
    fmt.Println("Output with Unicode type:")
    printBytes(name)
    fmt.Println("Output with char type:")
    printChars(name)
    name = "你好, world!"
    fmt.Printf("\n")
    printChars(name)
    fmt.Printf("Print Unicode bytes length:\n")
    printCharsAndBytes(name)

    word1 := "Hello, 世界！" 
    length(word1)
    word2 := "Pets"
    length(word2)

    h := "hello"
    // fmt.Println(mutate(h))
    fmt.Println(mutate_slice([]rune(h)))
    
}
