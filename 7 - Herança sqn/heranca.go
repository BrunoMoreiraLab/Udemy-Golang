package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	altura    int
	sexo      string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Bruno", "Moreira", 169, "Masculino"}
	fmt.Println(p1)

	e1 := estudante{p1, "Analise e Desenvolvimento de Sistemas", "UNIP"}
	fmt.Println(e1)
}
