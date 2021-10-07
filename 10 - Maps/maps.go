package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Bruno",
		"Sobrenome": "Cruz",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"Campus": "Campus I",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Escorpião",
	}
	fmt.Println(usuario2)
}
