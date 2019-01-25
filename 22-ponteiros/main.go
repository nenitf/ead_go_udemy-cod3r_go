package main

import (
	"fmt"
)

func main() {
	fmt.Println("Para variaveis: usar '&' para apontamento")
	fmt.Println("Para ponteiros: usar '&' para apontamento e '*' para seguir literalmente o valor do apontamento")

	// atribuição implicita (por inferência)
	i1 := 1
	p1 := &i1
	ii1 := *p1
	i1++
	iii1 := *p1
	*p1++
	imprime(i1, ii1, iii1, p1)

	fmt.Printf(`

===== print =====
// primeiro inteiro
i => %v,
&i => %v,

// primeiro ponteiro
p => %v,
&p => %v,
*p => %v

// segundo inteiro
ii => %v
&ii => %v

// segundo inteiro
iii => %v
&iii => %v`, i1, &i1, p1, &p1, *p1, ii1, &ii1, iii1, &iii1)

	// atribuição explícita
	var i2 int
	var p2 *int
	var ii2 int
	var iii2 int

	i2 = 2
	p2 = &i2
	ii2 = *p2
	i2++
	iii2 = *p2
	*p2++

	imprime(i2, ii2, iii2, p2)
	fmt.Printf(`

===== print =====
// primeiro inteiro
i => %v,
&i => %v,

// primeiro ponteiro
p => %v,
&p => %v,
*p => %v

// segundo inteiro
ii => %v
&ii => %v

// segundo inteiro
iii => %v
&iii => %v`, i2, &i2, p2, &p2, *p2, ii2, &ii2, iii2, &iii2)
}

func imprime(i, ii, iii int, p *int) {
	fmt.Printf(`

===== função =====
// primeiro inteiro
i => %v,
&i => %v,

// primeiro ponteiro
p => %v,
&p => %v,
*p => %v

// segundo inteiro
ii => %v
&ii => %v

// segundo inteiro
iii => %v
&iii => %v`, i, &i, p, &p, *p, ii, &ii, iii, &iii)

}
