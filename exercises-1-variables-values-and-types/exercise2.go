package ex1

import "fmt"

// Use var para declarar três variáveis.
// Elas devem ter package-level scope. Não atribua valores a estas variáveis.
// a deverá ser int, b deverá ser string e c deverá ser bool.

var a int
var b string
var c bool

// Print type of a variable
// https://gosamples.dev/print-type/

func ex2() {
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)
}
