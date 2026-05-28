package main

import "fmt"

func main() {

	var price float64 = 24.50;

	var mango_price float64 = 15.75;

	var pi float64 = 3.14159;

	var apple_price = 30.00;

	var orange_price = 20.00;


	var total_price = price + mango_price + apple_price + orange_price;
	 
	fmt.Println("Price:", price)
	fmt.Println("Mango Price:", mango_price)
	fmt.Println("Pi:", pi)

	fmt.Println("Apple Price:", apple_price)
	fmt.Println("Orange Price:", orange_price)
	fmt.Println("Total Price:", total_price)

}