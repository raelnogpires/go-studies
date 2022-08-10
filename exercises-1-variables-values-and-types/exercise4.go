package ex1

import "fmt"

// Crie um tipo. O tipo subjacente deve ser int.
// Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.

type w int

var h w

func ex4() {
	fmt.Println(h)
	fmt.Printf("h: %T\n", h)

	h = 10
	fmt.Println(h)
}
