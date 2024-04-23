package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoID}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoID}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoID}/editar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
