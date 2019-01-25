package main

import "fmt"

func main() {
	// constantes
	// palavra reservada const +
	// nome da constante +
	// tipo
	const pi float64 = 3.1415

	// bloco
	const (
		a  = 1
		aa = 11
	)

	// variaveis
	// palavra reservada var +
	// nome
	var raio = 10

	// bloco, inline e inline inferido
	var (
		z = 1
		x = 2
	)
	q, w := 1, 2
	var e, r bool = true, false

	// forma reduzida
	// inferida pela linguagem
	area := 3.2
	_ = "sem problema"

	// todas variaveis e constantes declaradas devem ser utilizadas
	// a menos que seja usado o blank indentifier _
	fmt.Println(pi)
	fmt.Println(raio)
	fmt.Println(area)
	fmt.Println(a)
	fmt.Println(q, w)
	fmt.Println(z, x)
	fmt.Println(aa)
	fmt.Println(e, r)
}
