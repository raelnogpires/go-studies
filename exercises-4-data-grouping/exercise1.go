package main

import "fmt"

// create an array that supports 5 int values
// each index receive a value
// using `range`, print each value
// using format printing, print the type of the array

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for _, i := range arr {
		fmt.Println(i)
	}

	fmt.Printf("%T\n", arr)
}
