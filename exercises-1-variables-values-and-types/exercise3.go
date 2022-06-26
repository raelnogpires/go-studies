package main

import "fmt"

// Em package-level scope, atribua os seguintes valores às variáveis:
// x recebe 42, y recebe "James Bond" e z recebe true.

var x int = 42
var y string = "James Bond"
var z bool = true

// Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
// Faça essa atribuição de tipo string a uma variável de nome "s"
// utilizando o operador curto de declaração.

func main() {
	s := fmt.Sprintf("%v, %v, %v", x, y, z)

	fmt.Println(s)
}
