package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

//rotasUsuarios, o metodo de cada rota a ser utilizada
var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios", // CADASTRA USUARIOS
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{
		Uri:                "/usuarios", // BUSCA USUARIOS
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},

	{
		Uri:                "/usuarios/{usuarioId}", //BUSCA USUARIO
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},

	{
		Uri:                "/usuarios/{usuarioId}", // ATUALIZAR USUARIO
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},

	{
		Uri:                "/usuarios/{usuarioId}", // DELETAR USUARIO
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
