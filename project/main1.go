package main

import (
    "fmt"
)

func main() {
    var a, b int
    var operator string

    fmt.Println("===== Calculator CLI =====")

    fmt.Print("Enter first number: ")
    fmt.Scan(&a)

    fmt.Print("Enter operator (+ - * /): ")
    fmt.Scan(&operator)

    fmt.Print("Enter second number: ")
    fmt.Scan(&b)

    switch operator {

    case "+":
        fmt.Println("Result =", a+b)

    case "-":
        fmt.Println("Result =", a-b)

    case "*":
        fmt.Println("Result =", a*b)

    case "/":
        if b != 0 {
            fmt.Println("Result =", a/b)
        } else {
            fmt.Println("Division by zero not allowed")
        }

    default:
        fmt.Println("Invalid Operator")
    }
}