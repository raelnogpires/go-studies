package main

import "fmt"

// Crie um tipo. O tipo subjacente deve ser int.
// Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.

type something int

var x something

func main() {
	fmt.Println(x)
	fmt.Printf("x: %T\n", x)

	x = 10
	fmt.Println(x)
}
