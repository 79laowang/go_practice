package main

import(
    "fmt"
)

func main() {
    a := [5]int{3, 4, 76, 79, 80}
    var b []int = a[1:4] // creates a slice from a[1] to a[3]
    fmt.Println(b)

    fmt.Println("Init a slice:")
    c := []int{6, 7, 8}  //creates and array and returns a slice reference
    fmt.Println(c)

    fmt.Println("Modify array by slice changes:")
    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before:", darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after:", darr)

    fmt.Println("Multiple slices changes on array:")
    numa := [3]int{78, 79, 80}
    nums1 := numa[:]
    nums2 := numa[:]
    fmt.Println("array before change:", numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1:", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2:", numa)

    /*
    The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
    */
    fmt.Println("\nlength and capacity of a slice:")
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
    fmt.Println("\nfruitslice:",fruitslice)
    fruitslice = fruitslice[:cap(fruitslice)] // re-slicing fruitslice till its capacity
    fmt.Printf("After re-slcing length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
    fmt.Println("\nfruitslice:",fruitslice)

    fmt.Println("\ncreating a slice using make:")
    islice := make([]int, 5, 5)
    fmt.Println(islice)

    fmt.Println("\nAppend a new item to slice:")
    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
    cars = append(cars, "Toyata")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
}

