package main

import "fmt"

func main() {

	var a int = 100
	var b int = 50

	if a > b {
		fmt.Println(a, "is largest")
	} else {
		fmt.Println(b, "is largest")
	}
}