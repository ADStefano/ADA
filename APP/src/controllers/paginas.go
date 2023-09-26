package controllers

import (
	"app/utils"
	"net/http"
)

// Renderiza a tela de login
func CarregarLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// Renderiza a tela de cadastro
func CarregarCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// Renderiza a pagina principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "home.html", nil)
}
