package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI: "/criar-usuario",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarCadastroUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CadastrarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/buscar-usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioID}",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioID}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioID}/parar-de-seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
}
