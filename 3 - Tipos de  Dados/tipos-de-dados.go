package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	var numero int64 = -100000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias = Apelidos
	//int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	//Byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMEROS REAIS FLOAT32, FLOAT34

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12345.45
	fmt.Println(numeroReal2)

	numeroReal3 := 98765.98
	fmt.Println(numeroReal3)

	// STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// VALOR ZERO
	var zero int16
	fmt.Println(zero)

	var texto string
	fmt.Println(texto)

	//BOLEANO

	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano := true
	fmt.Println(booleano)

	//ERROS

	var erroo error
	fmt.Println(erroo)

	var erro = errors.New("Erro interno")
	fmt.Println(erro)
}
