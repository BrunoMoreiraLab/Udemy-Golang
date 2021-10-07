package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) Aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Bruno", 20}
	usuario1.salvar()

	usuario2 := usuario{"Vanessa", 20}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.Aniversario()
	fmt.Println(usuario2.idade)
}
