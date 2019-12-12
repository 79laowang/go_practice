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

    fmt.Println("Accessing items of a map:")
    workerSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    workerSalary["mike"] = 9000
    employee := "jamie"
    fmt.Println("Salary of", employee, "is", workerSalary[employee])
    // if we try to access an element which is not present then, the zero value of int which is 0 will be returned. 
    fmt.Println("Salary of joe is", workerSalary["joe"])

    fmt.Println("Checking whether a key is present in a map or not.")
    newEmp := "joe"
    value, ok := personSalary[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
    } else {
        fmt.Println(newEmp, "Not found")
    }

    fmt.Println("All items of a map:")
    for key, value := range workerSalary {
        fmt.Printf("workerSalary[%s] = %d\n", key, value)
    }

    fmt.Println("Deleting a item from map:")
    delete(workerSalary, "mike")
    fmt.Println("map after deletion", workerSalary)

    fmt.Println("length of map is:", len(workerSalary))

    fmt.Println("Maps are reference types:")
    workerSalary["mike"] = 9000
    fmt.Println("Original worker salary:", workerSalary)
    newPersonSalary := workerSalary
    newPersonSalary["mike"] = 18000
    fmt.Println("Worker salary changed from reference map:", workerSalary)

    fmt.Println("Maps equality:")
    map1 := map[string]int{
        "one": 1,
        "two": 2,
    }
    var map2 map[string]int
    if map2 == nil {
        fmt.Println("map is empty!")
        map2 = map1
    }
    fmt.Println("map2:", map2, "map1:", map1)
}
