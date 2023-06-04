package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Adiciona uma nova publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r* http.Request){

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(requestBody, &publicacao); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)
	
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

// Traz as publicações que apareceriam no feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r* http.Request){

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)

	publicacoes, erro := repositorio.BuscarPublicacoes(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

// Traz uma única publicação
func BuscarPublicacao(w http.ResponseWriter, r* http.Request){

	parametro := mux.Vars(r)

	publicacaoID, erro := strconv.ParseUint(parametro["PublicacaoId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)

	publicacao, erro := repositorio.BuscarPorId(publicacaoID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)
}

// Traz as publicações de um usuário específico
func BuscarPublicacoesUsuario(w http.ResponseWriter, r* http.Request){

	parametro := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametro["UsuarioId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)

	publicacoes, erro := repositorio.BuscarPublicacoesUsuario(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

// Altera os dados  de uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r* http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["PublicacaoId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)

	publicacaoNoBanco, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoNoBanco.AutorID != usuarioID{
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
		return
	}

	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao

	if erro = json.Unmarshal(requestBody, &publicacao); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Preparar(); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicacaoId, publicacao); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError	, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// Deleta uma publicação
func DeletarPublicacao(w http.ResponseWriter, r* http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["PublicacaoId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)

	publicacaoNoBanco, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoNoBanco.AutorID != usuarioID{
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar uma publicação que não seja sua"))
		return
	}

	if erro = repositorio.Deletar(publicacaoId); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError	, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// Adiciona uma curtida a uma publicação no banco de dados
func CurtirPublicacao(w http.ResponseWriter, r* http.Request){

	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["PublicacaoId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)

	if erro = repositorio.Curtir(publicacaoId); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}

// Remove uma curtida de uma publicação no banco de dados
func DescurtirPublicacao(w http.ResponseWriter, r* http.Request){

	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["PublicacaoId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conn()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioPublicacoes(db)

	if erro = repositorio.Descurtir(publicacaoId); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
