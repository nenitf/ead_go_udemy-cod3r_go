package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)
	fmt.Println(2 > 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 < 3)
	fmt.Println(2 <= 3)

	// datas
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("datas:", d1 == d1)     // true
	fmt.Println("datas:", d1.Equal(d2)) // true

	// structs
	type Pessoa struct {
		Nome string
	}
	p1 := Pessoa{"joão"}
	p2 := Pessoa{"joão"}
	fmt.Println("Pessoas:", p1 == p2) // true
}
