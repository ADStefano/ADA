package controllers

import (
	"app/utils"
	"net/http"
)

// Renderiza a tela de login
func CarregarLogin(w http.ResponseWriter, r *http.Request){
	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarCadastroUsuario(w http.ResponseWriter, r *http.Request){
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}