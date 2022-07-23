package main

import "fmt"

// create typed and untyped constants
// print the values

const typed int = 42
const untyped = "this variable has no type"

func main() {
	fmt.Println(typed)
	fmt.Println(untyped)
}
