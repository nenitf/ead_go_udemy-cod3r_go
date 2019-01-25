package main

func somar(ns ...int) int {
	var resposta int
	for _, num := range ns {
		resposta += num
	}
	return resposta
}
