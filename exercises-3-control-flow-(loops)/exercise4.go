package main

import "fmt"

// print the module of divison by 4 of all number from 10 to 100.

func main() {
	for i := 10; i <= 100; i += 1 {
		fmt.Println(i % 4)
	}
}
