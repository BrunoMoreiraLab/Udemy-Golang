package main

import "fmt"

// FUNÇÃO PANIC É ACABAR COM A EXECUÇÃO ASSIM QUE A FUNÇÃO ESCOLHIDA FOR EXECUTADA
// FUNÇÃO RECOVER RECUPERA UMA FUNÇÃO QUE ESTA ENTRANDO EM PANIC
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com Sucesso! ")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6! ")
}

func main() {

	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução !")
}
