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
	{
		URI: "/perfil",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPerfilDoUsuarioLogado,
		RequerAutenticacao: true,
	},
	{
		URI: "/editar-usuario",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaEdicaoPerfilDoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/editar-usuario",
		Metodo: http.MethodPut,
		Funcao: controllers.EditarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
}
