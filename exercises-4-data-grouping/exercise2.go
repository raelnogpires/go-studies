package ex4

import "fmt"

// create a type int slice
// the slice receives 10 int values
// use range method to demonstrate all methods
// use format sprinting to print the type of the slice

func ex2() {
	// creating the slice
	slice := make([]int, 10)
	// assigning values to the slice
	for i := range slice {
		slice[i] = i
	}

	fmt.Printf("%T\n", slice)
}
