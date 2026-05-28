package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	fmt.Println("Before Swap")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a, b = b, a

	fmt.Println("After Swap")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}