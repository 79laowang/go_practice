package main

import(
    "fmt"
    "rectangle"
    "log"
)

/*
* package variables
* Build rectangle package first:
* go build ../rectangle/rectprops.go
*/
var rectLen, rectWidth float64 = -6, 7

/*
* init function to check if length and width are greater than zero
*/
func init(){
    println("main package initialized.")
    if rectLen < 0 {
        log.Fatal("length is less than zero.")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero.")
    }
}


/* imports the rectangle package and uses the Area and Diagonal function of it to
*   find the area and diagonal of the rectangle. 
*/
func main(){
    fmt.Println("Geometrical shape properties")
    fmt.Printf("area of rectangle: %.2f\n", rectangle.Area(rectLen, rectWidth))
    fmt.Printf("diagonal of rectangle: %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
