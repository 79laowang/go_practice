package main

import(
    "fmt"
)

func changeLocal(num [5]int){
    num[0] = 55
    fmt.Println("inside function:", num)
}

func main(){
    var a[3]int
    fmt.Println(a)
    a[0] = 1
    a[1] = 10
    a[2] = 100
    fmt.Println(a)
    b := [3]int{11, 22, 33}
    fmt.Println(b)
    c := [3]int{11}
    fmt.Println(c)
    d := [...]int{12, 24, 36}
    fmt.Println(d)

    a1 := [...]string{"USA", "China", "India", "Germany", "France"}
    b1 := a1
    b1[0] = "Singapore"
    fmt.Println("a1 is", a1)
    fmt.Println("b1 is", b1)

    num :=[...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function:", num)
    changeLocal(num)
    fmt.Println("after passing to function:", num)

    num1 :=[...]float64{67.7, 11.1, 21, 78}
    for i:= 0; i < len(num1); i++ {
        fmt.Printf("%d th element of a is %.2f\n", i, num1[i])
    }

    num2 := [...]float64{67.7, 11.1, 21, 78}
    sum := float64(0)
    for j, v := range num2 {
        fmt.Printf("%d the element of a is %.2f\n", j, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of num2", sum)
}
