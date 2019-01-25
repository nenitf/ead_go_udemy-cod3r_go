package main

import "fmt"

func main() {
	imprimeResultado(3.2)
	imprimeResultado(9.9)
}

func imprimeResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com a nota", nota)
	} else {
		fmt.Println("Reprovado com a nota", nota)
	}
}
