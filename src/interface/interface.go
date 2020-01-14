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
    fmt.Printf("Type = %T, value = %v\n", i, i)
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
}
