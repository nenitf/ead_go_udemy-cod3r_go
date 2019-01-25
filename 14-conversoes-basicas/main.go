package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// float
	x := 2.4

	// int
	y := 2

	// não funciona
	// fmt.Println(x / y)

	fmt.Println(x / float64(y))
	float := 6.40949
	fmt.Println(int(float)) // imprime somente 6
	// fmt.Println(string(97)) // 97 decimal na tabelas ascii é "a"

	// inteiro para string
	fmt.Println(strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("1111")
	fmt.Println("num:", num, reflect.TypeOf(num))

	// string para bool
	b1, _ := strconv.ParseBool("true")     // true
	b2, _ := strconv.ParseBool("false")    // false
	b3, _ := strconv.ParseBool("ushdsujd") // false
	b4, _ := strconv.ParseBool("1")        // true
	b5, _ := strconv.ParseBool("0")        // false
	b6, _ := strconv.ParseBool("-123")     // false
	b7, _ := strconv.ParseBool("123")      // false

	fmt.Println("b:", b1, reflect.TypeOf(b1))
	fmt.Println("b:", b2, reflect.TypeOf(b2))
	fmt.Println("b:", b3, reflect.TypeOf(b3))
	fmt.Println("b:", b4, reflect.TypeOf(b4))
	fmt.Println("b:", b5, reflect.TypeOf(b5))
	fmt.Println("b:", b6, reflect.TypeOf(b6))
	fmt.Println("b:", b7, reflect.TypeOf(b7))
}
