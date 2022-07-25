package main

import "fmt"

// create a program that demonstrates the functionality of the if statement.

func testIf(arg bool) {
	if arg {
		fmt.Println("true :)")
	} else {
		fmt.Println("false :(")
	}
}

func main() {
	testIf(true)
	testIf(false)
}
