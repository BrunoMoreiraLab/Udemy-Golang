package main

import "fmt"

func main() {

	// 1 - EXEMPLO
	func() {
		fmt.Println("Olá Mundo")
	}()

	// 2 - EXEMPLO
	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	// 3 - EXEMPLO
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)

	// 4 - EXEMPLO
	retorno2 := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s, %d", texto, 10)
	}("Passando Parâmetro")

	fmt.Println(retorno2)
}
