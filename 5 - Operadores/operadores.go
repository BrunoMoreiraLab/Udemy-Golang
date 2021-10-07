package main

import "fmt"

func main() {

	// ARITMETICOS
	fmt.Println("-----------------")
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	somar := numero1 + numero2
	fmt.Println(somar)
	fmt.Println("-----------------")
	//FIM ARITMETICO

	// ATRIBUIÇÃO
	fmt.Println("-----------------")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	fmt.Println("-----------------")
	// FIM ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println("-----------------")
	fmt.Println(1 > 2)  //Maior
	fmt.Println(1 < 2)  //Menor
	fmt.Println(1 >= 2) //Menor Igual
	fmt.Println(1 <= 2) //Menor igual
	fmt.Println(1 == 2) //Comparação
	fmt.Println(1 != 2) //Diferença
	fmt.Println("-----------------")
	// FIM RELACIONAIS

	// OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println("-----------------")
	fmt.Println(!falso && verdadeiro)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	fmt.Println("-----------------")
	//FIM OPERADORES LÓGICOS

	// OPERADORES UNARIOS
	fmt.Println("-----------------")
	Incremento := 10
	Incremento++
	Incremento += 15 // Incremento = Incremento + 15
	fmt.Println(Incremento)

	Decremento := 20
	Decremento--
	Decremento -= 15 // Decremento - 15
	fmt.Println(Decremento)
	//FIM OPERADORES UNARIOS
}
