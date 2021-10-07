package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	// 1 - EXEMPLO
	numero := 10

	if numero > 15 {
		fmt.Println(numero, "É maior que 15")
	} else {
		fmt.Println(numero, "É menor ou igual a 15")
	}

	// 2 - EXEMPLO

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println(numero, "Numero é maior que zero")
	} else if numero < -10 {
		fmt.Println(numero, "Numero é menor que zero")
	} else {
		fmt.Println("Entre 0 e - 10")
	}
}
