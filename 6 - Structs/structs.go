package main

import (
	"fmt"
)

type usuario struct {
	nome      string
	sobrenome string
	idade     int
	endereco  endereco
}

type endereco struct {
	logradouro string
	cidade     string
	numero     int
}

func main() {
	fmt.Println("ARQUIVO STRUCT")

	// 1- JEITO DE SER FEITO
	var u usuario
	u.nome = "Bruno"
	u.sobrenome = "Moreira"
	u.idade = 25
	fmt.Println(u)

	//STRUCT DENTRO DE STRUCT
	enderecoExemplo := endereco{"Av. Milton de Oliveira", "Praia Grande", 445}

	// 2- JEITO DE SER FEITO
	usuario2 := usuario{"Vanessa", "Cruz", 21, enderecoExemplo}
	fmt.Println(usuario2)

}
