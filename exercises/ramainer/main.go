package main

import (
	"fmt"
)

func main() {
	var smallNumber int
	var largerNumber int

	fmt.Println("Insert a small numer")
	fmt.Scan(&smallNumber)

	fmt.Println("Insert a larger number")
	fmt.Scan(&largerNumber)

	fmt.Printf("The remainer of %d and %d is: %d \n", smallNumber, largerNumber, largerNumber%smallNumber)
}
