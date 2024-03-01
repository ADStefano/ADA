package rotas

import (
	"app/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Representa todas as rotas da aplicação web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		
		if rota.RequerAutenticacao{router.HandleFunc(rota.URI, middleware.Logger(middleware.Autenticar(rota.Funcao))).Methods(rota.Metodo)
			
		} else {
			router.HandleFunc(rota.URI, middleware.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
