package ex3

import "fmt"

// create a program that demonstrates the functionality of the if statement.

func testIf(arg bool) {
	if arg {
		fmt.Println("true :)")
	} else {
		fmt.Println("false :(")
	}
}

func ex5() {
	testIf(true)
	testIf(false)
}
