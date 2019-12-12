package main

import(
    "fmt"
)

/*variadic function
* The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter.
*/
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v :=range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "Not found in ", nums)
    }
    fmt.Printf("\n")
}

func main(){
    find(89, 89, 90, 95)
    find(45, 56, 45, 67, 90, 109)
    find(78, 38, 56, 98)
    find(85)
}
