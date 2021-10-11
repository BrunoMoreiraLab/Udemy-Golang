package rotas

import (
	"api/src/controllers"
	"net/http"
)

//rotasUsuarios, o metodo de cada rota a ser utilizada
var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios", // CADASTRA USUARIOS
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios", // BUSCA USUARIOS
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}", //BUSCA USUARIO
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}", // ATUALIZAR USUARIO
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}", // DELETAR USUARIO
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
