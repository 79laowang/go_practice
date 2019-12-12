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

/* Using slice to implement the same function as variadic */
func findInSlice(num int, nums []int){
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "Found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "Not found in", nums)
    }
    fmt.Printf("\n")
}

func change(s ...string) {
    s[0] = "Go"
    s = append(s, "playground")
    fmt.Println("append element to slice s", s)
}

func main(){
    fmt.Println("Found with variadic:")
    find(89, 89, 90, 95)
    find(45, 56, 45, 67, 90, 109)
    find(78, 38, 56, 98)
    find(85)

    fmt.Println("Found with slice:")
    findInSlice(89, []int{89, 90, 95})
    findInSlice(45, []int{56, 45, 67, 90, 109})
    findInSlice(78, []int{38, 56, 98})
    findInSlice(85, []int{})

    fmt.Println("Found with slice in variadic function :")
    nums := []int{89, 90, 100}
    //syntactic sugar which can be used to pass a slice to a variadic function
    find(100, nums...)

    fmt.Println("Change slice inside a variadic function:")
    welcome := []string{"hello", "world"}
    fmt.Println("Original string:", welcome)
    change(welcome...)
    fmt.Println("After changed in variadic function:", welcome)
}
