package ex1

import "fmt"

// Em package-level scope, utilizando a palavra-chave var, crie uma variável com
// o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você
// criou no exercício anterior.

var y int

type something int

var x something

func ex5() {
	fmt.Println(x)
	fmt.Printf("x: %T\n", x)

	x = 10
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("y: %T\n", y)
}
