package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma:", a+b)
	fmt.Println("subtração:", a-b)
	fmt.Println("divisão:", a/b)
	fmt.Println("multiplicação:", a*b)
	fmt.Println("módulo:", a%b)

	// bitwise
	// binários
	// a: 3 => 11
	// b: 2 => 10
	fmt.Println("E:", a&b)   // 11 & 10 = 10 => 2
	fmt.Println("OU:", a|b)  // 11 | 10 = 11 => 3
	fmt.Println("XOR:", a^b) // 11 ^ 10 = 01 => 1

	// precisa ser float64
	fmt.Println("Maior:", math.Max(float64(a), float64(b))) // precisa ser float64
}
