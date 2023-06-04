package rotas

import (
	"api/src/controllers"
	"net/http"
)

var RotasUsuarios = []Rota{
	{
		URI : "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI : "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI : "/usuarios/{UsuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI : "/usuarios/{UsuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI : "/usuarios/{UsuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{UsuarioId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{UsuarioId}/parar-seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{UsuarioId}/seguidores",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{UsuarioId}/seguindo",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{UsuarioId}/atualizar-senha",
		Metodo: http.MethodPost,
		Funcao: controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}