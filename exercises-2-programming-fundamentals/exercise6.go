package ex2

import "fmt"

// using iota, create 4 constants which the values are the 4 next years
// print the results

const (
	year1 = 2023 + iota
	year2
	year3
	year4
)

func ex6() {
	fmt.Println(year1, year2, year3, year4)
}
