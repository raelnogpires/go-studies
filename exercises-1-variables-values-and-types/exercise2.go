package main

import "fmt"

// Use var para declarar três variáveis.
// Elas devem ter package-level scope. Não atribua valores a estas variáveis.
// x deverá ser int, y deverá ser string e z deverá ser bool.

var x int
var y string
var z bool

// Print type of a variable
// https://gosamples.dev/print-type/

func main() {
	fmt.Printf("x: %T\n", x)
	fmt.Printf("y: %T\n", y)
	fmt.Printf("z: %T\n", z)
}
