package controllers

import (
	"app/src/cookies"
	"net/http"
)

// Remove os dados de autenticação do browser do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {

	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)

}
