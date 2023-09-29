package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rotas{
	{
		URI:          "/usuarios",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CriarUsuarios,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarUsarios,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarUsario,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodPut,
		Funcao:       controllers.AtualizarUsuario,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodDelete,
		Funcao:       controllers.DeletarUsuario,
		Autenticacao: false,
	},
}
