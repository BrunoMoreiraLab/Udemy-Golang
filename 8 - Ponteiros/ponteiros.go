package main

import "fmt"

func main() {

	// PONTEIRO É UMA REFERENCIA DE MEMORIA

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro) //local de memoria

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // valor referenciado da memoria
}
