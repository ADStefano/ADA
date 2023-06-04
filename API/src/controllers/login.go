package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"strconv"

	// "fmt"
	"io"
	"net/http"
)

// Função resposável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	RequestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuarios
	if erro = json.Unmarshal(RequestBody, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)

	usuarioBanco, erro := repositorio.BuscarEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	if erro := seguranca.VerificarSenha(usuarioBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var TokenJSON modelos.Token

	TokenJSON.ID = strconv.FormatUint(usuarioBanco.ID, 10)
	TokenJSON.Token = token

	respostas.JSON(w, http.StatusOK, TokenJSON)
}
