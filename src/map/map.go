package main

import(
    "fmt"
)

func main() {
    var personSalary map[string]int
    if personSalary == nil {
        fmt.Println("map is nil. Going to make one.")
        personSalary = make(map[string]int)
        personSalary["Bob"] = 10000
        personSalary["Lili"] = 18000
        personSalary["Ke"] = 30000
        fmt.Println("personSalary map contents:", personSalary)

    }
}
