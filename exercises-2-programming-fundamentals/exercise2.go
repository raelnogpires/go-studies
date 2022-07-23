package main

import "fmt"

// write expressions with the following operators and print the values:
// ==, !=, <=, <

var equal = (42 == 50)          // should log false
var diff = (42 != 50)           // should log true
var lessOrEqual = (42 <= 50)    // should log true
var less = (74 < 50)            // should log false
var greaterOrEqual = (42 >= 50) // should log false
var greater = (74 > 50)         // should log true

func main() {
	fmt.Println(equal)
	fmt.Println(diff)
	fmt.Println(lessOrEqual)
	fmt.Println(less)
	fmt.Println(greaterOrEqual)
	fmt.Println(greater)
}
