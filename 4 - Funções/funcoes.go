package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n3, n4 int) (int, int) {
	soma := n3 + n4
	subtracao := n3 - n4
	return soma, subtracao
}

func main() {
	soma := somar(19, 29)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(20, 20) // _ Substitui uma variavel
	fmt.Println(resultadoSoma)
}
