package main

import "fmt"

// DEFER TEM A FUNÇÃO DE ADIAR A EXECUÇÃO QUE FOI ATRIBUIDO
func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função e verificando se o aluno foi aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}
func main() {
	fmt.Println(alunoAprovado(7, 8))
}
