package main

// 	NÃO EXISTEEEE

import "fmt"

func main() {
	resposta := obterResultado(7)
	fmt.Println(resposta)
}

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	/* se fosse como js:
	return nota >= 6 ? "Aprovado" : "Reprovado"
	*/
}
