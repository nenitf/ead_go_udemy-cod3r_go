package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// inteiros
	fmt.Println(reflect.TypeOf(223))

	// somente positivos uint8, uint16, uint32, uint64
	// uint8 equivale a byte
	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	// com sinal int8, int16, int32, int64
	i1 := math.MaxInt32
	//i1 := math.MaxInt64 funciona em maquinas x64
	fmt.Println("O valor maximo de i1 é", i1)
	fmt.Println("o tipo i1 é", i1)

	var i2 rune = 'a' // tabela unicod (int32)
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2)

	// numero reais float32, float64
	f1 := 32.2222
	var f2 float32 = 43.44
	fmt.Println(f1, f2)

	// strings
	// não se usam aspas simples
	s1 := "uma string"
	fmt.Println(s1 + "!") // concatenar string
	fmt.Println("tamanho da string é", len(s1))

	s2 := `a
sd
	f
	f
	f
	f
	v
	v
	vv
	v
	`
	fmt.Println(s2)
}
