package main

import (
	"fmt"
)

func main() {

	/*// 1 EXEMPLO
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second) // Trabalha com o tempo e data

	}
	fmt.Println(i)*/

	/*// 2 EXEMPLO
	for j := 0; j <= 10; j += 2 {
		fmt.Println("Incremento j", j)
		time.Sleep(time.Second)
	}*/

	// 3 - EXEMPLO

	nomes := [3]string{"Bruno", "Moreira", "Cruz"}

	for _, nome := range nomes {
		fmt.Println(nome)

	}

	// 4 - EXEMPLO

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// 	5 EXEMPLO
	usuario := map[string]string{
		"Nome":      "Bruno",
		"Sobrenome": "Cruz",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
