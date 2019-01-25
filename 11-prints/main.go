package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print("linha")
	fmt.Println("quebra proxima linha")

	x := 1.23456

	fmt.Println("o valor de x é", x) // espaço adicionado
	xs := fmt.Sprint(x)
	fmt.Println("concatenado-sem-espaço-" + xs)

	inteiro := 1
	flutuante := 1.23456
	boleano := true
	palavra := "qqqqq"

	fmt.Printf("\n%d %f %.2f %t %s", inteiro, flutuante, flutuante, boleano, palavra)
	fmt.Printf("\n%v %v %.2v %v %v", inteiro, flutuante, flutuante, boleano, palavra)

}
