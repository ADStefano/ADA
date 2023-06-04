package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaPublicacoes  = []Rota {

	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{PublicacaoId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{PublicacaoId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{PublicacaoId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{UsuarioId}/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacoesUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{PublicacaoId}/curtir",
		Metodo: http.MethodPost,
		Funcao: controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{PublicacaoId}/descurtir",
		Metodo: http.MethodPost,
		Funcao: controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}