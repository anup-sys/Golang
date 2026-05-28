package main

import (
    "fmt"
)

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
)

func main() {

    var a, b float64
    var operator string

    fmt.Println(Cyan + "============================")
    fmt.Println("     COLORFUL CALCULATOR")
    fmt.Println("============================" + Reset)

    fmt.Print(Yellow + "Enter first number: " + Reset)
    fmt.Scan(&a)

    fmt.Print(Yellow + "Enter operator (+ - * /): " + Reset)
    fmt.Scan(&operator)

    fmt.Print(Yellow + "Enter second number: " + Reset)
    fmt.Scan(&b)

    switch operator {

    case "+":
        fmt.Println(Green, "Result =", a+b, Reset)

    case "-":
        fmt.Println(Green, "Result =", a-b, Reset)

    case "*":
        fmt.Println(Green, "Result =", a*b, Reset)

    case "/":
        if b != 0 {
            fmt.Println(Green, "Result =", a/b, Reset)
        } else {
            fmt.Println(Red, "Error: Division by zero!", Reset)
        }

    default:
        fmt.Println(Red, "Invalid Operator!", Reset)
    }
}