package main

import (
	"fmt"
	"strconv"
)

// variable receives int as value
// print the value as decimal, binary and hexadecimal
// deslocate the bytes of the variable 1 to the left and store this value in another variable
// print the value of the second variable as decimal, binary and hexadecimal.

func main() {
	const primary int64 = 28
	// other way to demonstrate the value:
	// fmt.Printf("%d\t%b\t%#x\n", primary, primary, primary)
	fmt.Println("primary decimal -> ", primary)
	fmt.Println("primary binary -> ", strconv.FormatInt(primary, 2))
	fmt.Println("primary hex -> ", fmt.Sprintf("%x", primary))

	const secondary int64 = primary << 1
	// fmt.Printf("%d\t%b\t%#x\n", secondary, secondary, secondary)
	fmt.Println("secondary decimal -> ", secondary)
	fmt.Println("secondary binary -> ", strconv.FormatInt(secondary, 2))
	fmt.Println("secondary hex -> ", fmt.Sprintf("%x", secondary))
}
