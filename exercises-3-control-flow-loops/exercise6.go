package ex3

import "fmt"

// using the solution of exercise5, add the `else if` and `else`

func ifTest(fruit string) {
	if fruit == "apple" {
		fmt.Println("apples are red")
	} else if fruit == "blueberry" {
		fmt.Println("blueberry are blue")
	} else {
		fmt.Println("idk this fruit")
	}
}

func ex6() {
	ifTest("apple")
	ifTest("blueberry")
	ifTest("watermelon")
}
