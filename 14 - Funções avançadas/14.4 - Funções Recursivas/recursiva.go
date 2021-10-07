package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci((posicao - 1))
}

func main() {
	fmt.Println("FUNÇÕES RECURSIVAS")

	//Sequência Fibonacci: 1 1 2 3 5 8 13

	posicao := uint(15)

	for i := uint(0); i < posicao; i++ { // Para imprimir todo o fibonacci
		fmt.Println(fibonacci(i)) // Para imprimir todo o fibonacci
	}
	//fmt.Println(fibonacci(posicao)) // Para imprimir apenas o fibonacci pedido
}
