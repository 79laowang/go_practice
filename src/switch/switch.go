package main

import "fmt"

func number()int {
    num := 15 * 6
    return num
}


func main(){
    finger := 4
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default:
        fmt.Println("incorrect finger number")
    }

    // Multiple expressions in case
    letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u":
        fmt.Println("vowel.")
    default:
        fmt.Println("Not a vowel!")
    }

    // Expressionless switch
    num := 75
    switch {  // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Println("Number is less then 50 and greater than 0.")
    case num >= 51 && num <= 100:
        fmt.Println("Number is greater then 51 and less then 100.")
    case num >= 101:
        fmt.Println("Numer is greater than 101.")
    }

    // Fallthrough, cannot fallthrough final case in switch
    switch num = number(); { // num is not a constant, get value from func
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200\n", num)
    }
}
