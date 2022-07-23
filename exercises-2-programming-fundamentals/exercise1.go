package main

import (
	"fmt"
	"strconv"
)

// write a program that prints a number in decimal, binary and hexadecimal
func main() {
	decimal := 42
	fmt.Println("decimal number -> ", decimal)

	binary := int64(42)
	fmt.Println("binary number -> ", strconv.FormatInt(binary, 2))

	hexadecimal := fmt.Sprintf("%x", 42)
	fmt.Println("hex number -> ", hexadecimal)
}
