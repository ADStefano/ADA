package middleware

import (
	"app/src/cookies"
	"log"
	"net/http"
)

// Escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Método: %s, Request: %s, Host: %s\n", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autentica e verifica a existência de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil{
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return 
		}
		proximaFuncao(w, r)
	}
}
