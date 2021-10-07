package main

import "fmt"

// FUNÇÃO PONTEIRO

func inverterSinal(numero int) int {
	fmt.Println("Número invertido")
	return numero * -1

}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println("Número para inversão ", numero)

	novoNumero := 40
	fmt.Println("Número para inversão ", novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println("Numero invertido", novoNumero)
}
