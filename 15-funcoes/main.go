package main

import "fmt"

func main() {
	// a func está em outro arquivo,
	// porém ainda no pacote main
	r := somar(1, 2, 3)
	fmt.Println(r)
}
