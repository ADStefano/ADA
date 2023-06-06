package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Faz a requisição para criar o usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuarios

	if erro = json.Unmarshal(requestBody, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Preparar("cadastro"); erro != nil {
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

	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		fmt.Println(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// Busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := banco.Conn()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)

	usuario, erro := repositorio.BuscarID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

// Busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	nome_nick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conn()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)

	usuarios, erro := repositorio.Buscar(nome_nick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// Atualiza as informações de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return
	}

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

	if erro := usuario.Preparar("atualizacao"); erro != nil {
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

	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// Deleta um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar um usuário que não seja o seu"))
		return
	}

	db, erro := banco.Conn()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)
	if erro := repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// Adiciona o usuário que fez a requisição como seguidor do usuário passado na URI
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID == seguidorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, erro := banco.Conn()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)
	if erro := repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// Remove o usuário que fez a requisição como seguidor do usuário passado na URI
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID == seguidorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível parar de seguir você mesmo, você está condenado a sua própria compania"))
		return
	}

	db, erro := banco.Conn()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)
	if erro := repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// Retorna todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
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
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

// Busca todos os seguidores de um usuário
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
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
	usuarios, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// Atualiza a senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {

	usuarioToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["UsuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioToken != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar a senha de outro usuário"))
		return
	}

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var senha modelos.Senha

	if erro := json.Unmarshal(requestBody, &senha); erro != nil {
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

	senhaAntiga, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := seguranca.VerificarSenha(senhaAntiga, senha.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com a senha salva no banco"))
		return
	}

	senhaHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorio.AtualizarSenha(usuarioID, string(senhaHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
