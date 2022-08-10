package ex1

import "fmt"

// Em package-level scope, atribua os seguintes valores às variáveis:
// first recebe 42, second recebe "James Bond" e third recebe true.

var first int = 42
var second string = "James Bond"
var third bool = true

// Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
// Faça essa atribuição de tipo string a uma variável de nome "s"
// utilizando o operador curto de declaração.

func ex3() {
	s := fmt.Sprintf("%v, %v, %v", first, second, third)

	fmt.Println(s)
}
