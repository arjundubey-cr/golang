package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 12

	var ptr = &myNumber
	var ptr2 *int = &myNumber

	*ptr = *ptr + 2
	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Value of pointer is", ptr2)
	fmt.Println("Value of pointer is", *ptr, *ptr2)
	fmt.Println("Value of number is", myNumber)
}
