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
}
