package router

import (
	"app/src/router/rotas"

	"github.com/gorilla/mux"
)

// Retorna um router com todas as rotas configuradas
func Gerar() *mux.Router{
	
	r:= mux.NewRouter()

	return rotas.Configurar(r)
}