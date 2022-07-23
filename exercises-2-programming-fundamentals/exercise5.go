package main

import "fmt"

// create a string variable using raw string literal
// print the result

func main() {
	const s string = `Hello,
										 World!`
	fmt.Println(s)
}
