package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "blabla"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

	// Invertendo valor de variaveis

	variavel5, variavel6 = "variavel 6", "variavel 5"
	fmt.Println(variavel5, variavel6)

	// Variavel constantes

	const constante1 string = "Constante 1"
	fmt.Println(constante1)
}
