package main

import (
    "fmt"
)

// interface definition
type VowelsFinder interface {
    FindVowels() []rune
}

type MyString string

// MyString which is receiver implements VowelsFinder
func (ms MyString) FindVowels() []rune { // This is a method
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

type SalaryCalculator interface {
    CalculateSalary() int
}

type Permanent struct {
    empId     int
    basicpay  int
    pf        int
}

type Contract struct {
    empId     int
    basicpay  int
}

//Salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
    return p.basicpay + p.pf
}

//Salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
    return c.basicpay
}

/*
total expense is calculated by interating though the SalaryCalculator slice and summing the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

type Tester interface {
    Test()
}

type MyFloat float64

func (m MyFloat) Test() {
    fmt.Println(m)
}

func describe (t Tester) {
    fmt.Printf("Interface type %T value %v\n", t, t)
}

func describe1(i interface{}) {
    fmt.Printf("interface Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
    v, ok := i.(int)  //get the underlying int value from i 
    fmt.Println(v, ok)
}

func findType(i interface{}) {
    switch i.(type) {
    case string:
        fmt.Printf("a string and value is %s\n", i.(string))
    case int:
        fmt.Printf("an int and value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}

type Describer interface {
    Describe()
}

type Person struct {
    name string
    age int
}

func (p Person) Describe() { //implemented using value receiver
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func findType1(i interface{}) {
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("Unknown type\n")
    }
}

type Address struct {
    state    string
    country  string
}

func (a *Address) Describe() { // implemeted using pointer receiver
    fmt.Printf("State %s Country %s\n", a.state, a.country)
}

type SalaryCalculator1 interface {
    DisplaySalary()
}

type LeaveCalculator interface {
    CalculateSalaryLeft() int
}

type Employee struct {
    firstName string
    lastName  string
    basicPay  int
    pf        int
    totalLeaves  int
    leavesTaken  int
}

func (e Employee) DisplaySalary() {
    fmt.Printf("%s %s has salary $%d", e.firstName , e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateSalaryLeft() int {
    return e.totalLeaves - e.leavesTaken
}

func main() {
    name := MyString("Sam Anderson")
    var v VowelsFinder
    // name is a receiver with a method FindVowels, so name is implemention  of
    // interface VowelsFinder
    v = name // possible since MyString implements VowelsFinder
    fmt.Printf("Vowels are %c", v.FindVowels())

    fmt.Println("\n\nPractical use of interface,")
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

    fmt.Println("\n\n")
    var t Tester
    f := MyFloat(89.7)
    t = f
    describe(t)
    t.Test()

/*
An interface which has zero methods is called empty interface. It is represented as interface{}. Since the empty interface has zero methods, all types implement the empty interface. 
*/
    fmt.Println("\n\nEmpty interface:")
    s := "Hello World"
    describe1(s)
    i := 55
    describe1(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe1(strt)

    /* i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.
    */
    fmt.Println("\n\nAssertion usage of interface:")
    var s1 interface{} = 56
    assert(s1)
    var i1 interface{} = "Steven Paul"
    assert(i1)

    fmt.Println("\n\ni.(T) usage for switch:")
    findType("Bob")
    findType("77")
    findType("89.88")

    findType1("Naveen")
    p1 := Person{
        name: "Naveen R",
        age:  25,
    }
    findType1(p1)

    fmt.Println("\n\nInterface using pointer receiver: ")
    var d1 Describer
    pp1 := Person{"Sam", 25}
    d1 = pp1
    d1.Describe()
    p2 := Person{"James", 32}
    d1 = &p2
    d1.Describe()
    var d2 Describer
    a1 := Address{"Washington", "USA"}
    d2 = &a1
    d2.Describe()

    fmt.Println("\n\nMultiple interfaces:")
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var s11 SalaryCalculator1 = e
    s11.DisplaySalary()
    var l LeaveCalculator = e
    fmt.Println("\nLeaves left =", l.CalculateSalaryLeft())

}
