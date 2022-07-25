package main

import "fmt"

// create a program that use switch statement without switch statement specified.

func main() {
	thirsty := 2

	switch {
	case thirsty == 1:
		fmt.Println("I want one beer, please.")
	case thirsty == 2:
		fmt.Println("Let's get a six-pack!")
	case thirsty == 3:
		fmt.Println("Oh boy, tonight is going to be a long night!")
	}
}
