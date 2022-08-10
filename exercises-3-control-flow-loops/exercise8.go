package ex3

import "fmt"

// create a program that use switch statement where the switch are
// a string variable with the identifier `favSport`

func ex8() {
	favSport := "soccer"

	switch favSport {
	case "soccer":
		fmt.Println("flamengo <3")
		break
	case "volley":
		fmt.Println("bruninho and wallace are unstoppable")
		break
	case "basketball":
		fmt.Println("have you heard about that stephen curry guy?")
		break
	default:
		fmt.Println("I don't know that sport!")
		break
	}
}
