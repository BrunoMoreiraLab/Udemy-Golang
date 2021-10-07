// HTTP É UM PROTOCOLO DE MOCUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)
//REQUEST - RESPONSE
// ROTAS
// URI - Identificador do recurso
// MÉTODO - GET, POST, PUT, DELETE

package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuario!"))
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/user", user)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
