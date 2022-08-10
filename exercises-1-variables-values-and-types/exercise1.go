package ex1

import "fmt"

// Utilizando o operador curto de declaração,
// atribua estes valores às variáveis com os identificadores "x", "y", e "z".
// I. 42, II. "James Bond", III. true

// Demonstre os valores dessas variáveis usando uma única declaração de print
// e também múltiplas declarações

func ex1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
